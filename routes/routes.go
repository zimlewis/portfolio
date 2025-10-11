package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zimlewis/portfolio/gintemplrenderer"
)


func InitRoutes(server *gin.Engine) {

	ginHtmlRenderer := server.HTMLRender

	server.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{
		FallbackHtmlRenderer: ginHtmlRenderer,
	}
	// htmlRenderer := server.HTMLRender
	// server.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{
	// 	FallbackHTMLRenderer: htmlRenderer,
	// }
	//
	//
	// server.SetTrustedProxies(nil)

	server.GET("/", home)
	server.GET("/projects", projects)
}
