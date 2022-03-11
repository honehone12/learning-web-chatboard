package route

import (
	"chatboard/common"
	"chatboard/models"
	"chatboard/user"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func getLogin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}

func getSignUp(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signup.html", nil)
}

func postSignUpAccount(ctx *gin.Context) {
	name := ctx.PostForm("name")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	if len(name) > 0 && len(email) > 0 && len(password) > 0 {
		res := user.CallService(&common.Message{
			Service:  common.ServiceCall,
			FuncType: user.CreateUser,
			Data: models.User{
				Name:     name,
				Email:    email,
				Password: password,
			},
		})
		if _, ok := res.Data.(error); ok {
			if gin.IsDebugging() {
				redirectToError(ctx, res.Data.(error).Error())
				return
			}
		} else {
			ctx.Redirect(http.StatusFound, "/user/login")
			return
		}
	}
	redirectToError(ctx, "sorry")
}

//////////////////////////////////////////////////
// is good way encrypting user data here??
func postAuthenticate(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	if len(email) > 0 && len(password) > 0 {
		res := user.CallService(&common.Message{
			Service:  common.ServiceCall,
			FuncType: user.GetUserByEmail,
			Data:     email,
		})
		if userAuth, ok := res.Data.(*models.User); ok {
			if strings.Compare(
				userAuth.Password,
				common.Encrypt(password),
			) == 0 {
				res = user.CallService(&common.Message{
					Service:  common.ServiceCall,
					FuncType: user.CreateSession,
					Data:     *userAuth,
				})
				if session, ok := res.Data.(*models.Session); ok {
					ctx.SetSameSite(http.SameSiteStrictMode)
					ctx.SetCookie(
						"short-time",
						session.UuId,
						0,
						"/",
						"localhost",
						true,
						true,
					)
					ctx.Redirect(http.StatusFound, "/")
					return
				} else if gin.IsDebugging() {
					redirectToError(ctx, res.Data.(error).Error())
					return
				}
			}
		} else if gin.IsDebugging() {
			redirectToError(ctx, res.Data.(error).Error())
			return
		}
	}
	ctx.Redirect(http.StatusFound, "/user/login")
}

func getLogOut(ctx *gin.Context) {
	uuid, err := ctx.Cookie("short-time")
	if err == nil {
		user.CallService(&common.Message{
			Service:  common.ServiceCall,
			FuncType: user.DeleteSessionByUUID,
			Data:     uuid,
		})
	}
	ctx.Redirect(http.StatusFound, "/")
}
