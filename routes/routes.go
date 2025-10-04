package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zimlewis/portfolio/gintemplrenderer"
)


func InitRoutes(server *gin.Engine) {
	htmlRenderer := server.HTMLRender
	server.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{
		FallbackHTMLRenderer: htmlRenderer,
	}


	server.SetTrustedProxies(nil)

}
