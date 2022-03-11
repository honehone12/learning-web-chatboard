package route

import (
	"chatboard/common"
	"chatboard/models"
	"chatboard/user"
	"fmt"
	"net/http"

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
		if affected, ok := res.Data.(int64); ok {
			if affected == 1 {
				ctx.Redirect(http.StatusFound, "/user/login")
				return
			} else if gin.IsDebugging() {
				redirectToError(ctx, fmt.Sprintf("returned value was %d", affected))
				return
			}
		} else if gin.IsDebugging() {
			redirectToError(ctx, res.Data.(error).Error())
			return
		}
	}
	redirectToError(ctx, "sorry")
}

func postAuthenticate(ctx *gin.Context) {

}

func getLogOut(ctx *gin.Context) {

}
