package route

import (
	"chatboard/common"
	"chatboard/models"
	"chatboard/thread"
	"chatboard/user"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const replyForm template.HTML = `<div class="panel panel-info">
  <div class="panel-body">
    <form id="post" role="form" action="/thread/post" method="post">
       <div class="form-group">
	     <textarea class="form-control" name="body" id="body" placeholder="Write your reply here" rows="3"></textarea>
	     <!-- get url with javascript? <input type="hidden" name="uuid" value=""> -->
		 <br/>
	     <button class="btn btn-primary pull-right" type="submit">Reply</button>
       </div>
    </form>
  </div>
</div>`

func getNewThread(ctx *gin.Context) {
	err := checkLoggedIn(ctx)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/user/login")
	} else {
		ctx.HTML(http.StatusFound, "newthread.html", gin.H{
			"navbar": privateNavbar,
		})
	}
}

func postNewThread(ctx *gin.Context) {
	err := checkLoggedIn(ctx)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/login")
	} else {
		uuid, _ := ctx.Cookie("short-time")
		res := user.CallService(&common.Message{
			Service:  common.ServiceCall,
			FuncType: user.GetSessionByUUID,
			Data:     uuid,
		})
		if sess, ok := res.Data.(*models.Session); ok {
			res = thread.CallService(&common.Message{
				Service:  common.ServiceCall,
				FuncType: thread.CreateThread,
				Data: common.Contribution{
					Content:  ctx.PostForm("topic"),
					UserID:   sess.UserId,
					UserName: sess.Name,
				},
			})
			if _, ok := res.Data.(error); ok {
				if gin.IsDebugging() {
					redirectToError(ctx, err.Error())
					return
				}

				redirectToError(ctx, "sorry")
				return
			}
		}
	}
	ctx.Redirect(http.StatusFound, "/")
}

func getThread(ctx *gin.Context) {
	uuid := ctx.Query("id")
	res := thread.CallService(&common.Message{
		Service:  common.ServiceCall,
		FuncType: thread.GetThreadByUUID,
		Data:     uuid,
	})
	if thre, ok := res.Data.(*models.Thread); ok {
		// can be middleware ??
		ctx.SetCookie(
			"long-time",
			uuid,
			60*60*24*365,
			"/",
			"localhost",
			true,
			true,
		)
		res := thread.CallService(&common.Message{
			Service:  common.ServiceCall,
			FuncType: thread.GetAllPostsInThread,
			Data:     thre.Id,
		})
		if posts, ok := res.Data.([]models.Post); ok {
			err := checkLoggedIn(ctx)

			if err != nil {
				ctx.HTML(http.StatusOK, "thread.html", gin.H{
					"navbar": publicNavbar,
					"thread": thre,
					"reply":  nil,
					"posts":  posts,
					"token":  "easy-token",
				})
			} else {
				ctx.HTML(http.StatusOK, "thread.html", gin.H{
					"navbar": privateNavbar,
					"thread": thre,
					"reply":  replyForm,
					"posts":  posts,
					"token":  "easy-token",
				})
			}
		} else {
			if gin.IsDebugging() {
				redirectToError(ctx, res.Data.(error).Error())
				return
			}
		}
	} else {
		if gin.IsDebugging() {
			redirectToError(ctx, res.Data.(error).Error())
			return
		}
	}
	redirectToError(ctx, "sorry")
}

func postPostToThread(ctx *gin.Context) {
	var res *common.Message
	if err := checkLoggedIn(ctx); err != nil {
		ctx.Redirect(http.StatusFound, "/user/login")
	} else {

		// token must be changed everytime and
		// we have to remember token.
		token := ctx.PostForm("token")
		if strings.Compare(token, "easy-token") != 0 {
			ctx.Abort()
			return
		}

		body := ctx.PostForm("body")
		sessUuId, _ := ctx.Cookie("short-time")
		threUuId := ctx.PostForm("uuid")
		threId, err := strconv.Atoi(ctx.PostForm("id"))
		if err != nil {
			if gin.IsDebugging() {
				redirectToError(ctx, res.Data.(error).Error())
			} else {
				redirectToError(ctx, "sorry")
			}
		}
		res = user.CallService(&common.Message{
			Service:  common.ServiceCall,
			FuncType: user.GetSessionByUUID,
			Data:     sessUuId,
		})
		if sess, ok := res.Data.(*models.Session); ok {
			res = thread.CallService(&common.Message{
				Service:  common.ServiceCall,
				FuncType: thread.CreatePost,
				Data: common.Contribution{
					ThreadID: uint(threId),
					Content:  body,
					UserID:   sess.UserId,
					UserName: sess.Name,
				},
			})
			if affected, ok := res.Data.(int64); ok && affected == 1 {
				url := fmt.Sprint("/thread/read?id=", threUuId)
				ctx.Redirect(http.StatusFound, url)
				return
			}
		}
	}
	if gin.IsDebugging() {
		redirectToError(ctx, res.Data.(error).Error())
	} else {
		redirectToError(ctx, "sorry")
	}
}
