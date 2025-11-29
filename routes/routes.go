package routes

import (
	"fmt"
	"os"

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

	publicDirectory := os.Getenv("PUBLIC_DIRECTORY")
	if publicDirectory == "" {
		fmt.Println("Cannot setup public directory")
		os.Exit(0)
	}

	server.Static("/static", publicDirectory)
	server.GET("/", home)
	server.GET("/projects", projects)
}
