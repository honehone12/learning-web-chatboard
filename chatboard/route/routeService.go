package route

import (
	"chatboard/templates"

	"github.com/gin-gonic/gin"
)

func OpenService(webEngine *gin.Engine) (groups []*gin.RouterGroup) {
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

}

func getIndex(ctx *gin.Context) {
	templates.RenderIndex(ctx)
}
