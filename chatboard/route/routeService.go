package route

import (
	"chatboard/common"
	"chatboard/models"
	"chatboard/thread"
	"chatboard/user"
	"html/template"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	publicNavbar template.HTML = `<div class="navbar navbar-default navbar-static-top" role="navigation">
  <div class="container">
    <div class="navbar-header">
      <a class="navbar-brand" href="/">KEIJIBAN</a>
    </div>
    <div class="nav navbar-nav navbar-right">
      <a href="/user/login">Login</a>
    </div>
  </div>
</div>`

	privateNavbar template.HTML = `<div class="navbar navbar-default navbar-static-top" role="navigation">
  <div class="container">
    <div class="navbar-header">
	  <a class="navbar-brand" href="/">KEIJIBAN</a>
    </div>
    <div class="nav navbar-nav navbar-right">
	  <a href="/user/logout">Logout</a>
    </div>
  </div>
</div>`
)

func OpenService(webEngine *gin.Engine) (groups []*gin.RouterGroup) {
	// setup templates
	webEngine.Static("/static", "./public")
	webEngine.Delims("{{", "}}")
	webEngine.LoadHTMLGlob("./templates/html/*")

	//setup routes
	webEngine.GET("/", getIndex)
	webEngine.GET(("/error"), getErr)

	userRoute := webEngine.Group("/user")
	userRoute.GET("/login", getLogin)
	userRoute.GET("/logout", getLogOut)
	userRoute.GET("/signup", getSignUp)
	userRoute.POST("/signup-account", postSignUpAccount)
	userRoute.POST("/authenticate", postAuthenticate)

	threadGroup := webEngine.Group("/thread")
	threadGroup.GET("/new", getNewThread)
	threadGroup.POST("/create", postNewThread)
	threadGroup.GET("/read", getThread)
	threadGroup.POST("/post", postPostToThread)
	return
}

func getErr(ctx *gin.Context) {
	errMsg := ctx.Query("msg")
	if err := checkLoggedIn(ctx); err != nil {
		ctx.HTML(http.StatusOK, "error.html", gin.H{
			"navbar": publicNavbar,
			"msg":    errMsg,
		})
	} else {
		//already logged in
		ctx.HTML(http.StatusOK, "error.html", gin.H{
			"navbar": privateNavbar,
			"msg":    errMsg,
		})
	}
}

func getIndex(ctx *gin.Context) {
	////////////////////////////////////////////////
	// here should be changed
	// get 10 or something
	res := thread.CallService(&common.Message{
		Service:  common.ServiceCall,
		FuncType: thread.GetAllThreads,
	})
	if threads, ok := res.Data.([]models.Thread); ok {
		if err := checkLoggedIn(ctx); err != nil {
			//already logged in
			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"navbar":  publicNavbar,
				"threads": threads,
			})
		} else {
			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"navbar":  privateNavbar,
				"threads": threads,
			})
		}
	} else {
		common.LogError().Println(res.Data.(error).Error())
		redirectToError(ctx, "could not get threads")
	}
}

func redirectToError(ctx *gin.Context, msg string) {
	url := []string{"/error?msg=", msg}
	ctx.Redirect(http.StatusFound, strings.Join(url, ""))
}

// this is login session
// i want this to be deleted on browser closed
func checkLoggedIn(ctx *gin.Context) (err error) {
	cookie, err := ctx.Cookie("short-time")
	if err == nil {
		res := user.CallService(&common.Message{
			Service:  common.ServiceCall,
			FuncType: user.GetSessionByUUID,
			Data:     cookie,
		})
		if _, ok := res.Data.(*models.Session); ok {
			// means just exist
		} else {
			err = res.Data.(error)
		}
	}
	return
}

// this is just a footprint
func checkFootprint(ctx gin.Context) {
	cookie, err := ctx.Cookie("long-time")
	if err == nil {
		// cookie is stored
		common.LogError().
			Printf("cookie %s found but not implemented yet.", cookie)
	} else {
		// cookie is notstored
		common.LogError().
			Printf("cookie not found and not implemented yet")
	}
}
