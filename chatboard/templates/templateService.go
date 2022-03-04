package templates

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OpenService(webEngine *gin.Engine) {
	webEngine.Static("/static", "./public")
	webEngine.Delims("{{", "}}")
	webEngine.LoadHTMLGlob("./templates/html/*")
}

func RenderErr(ctx *gin.Context) {

}

func RenderIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "layout.html", gin.H{})
}
