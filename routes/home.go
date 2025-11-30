package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zimlewis/portfolio/utils"
	"github.com/zimlewis/portfolio/views/homeview"
)


func home(c *gin.Context) {
	cookies := utils.GetCookies(c)
	c.HTML(http.StatusOK, "", homeview.Home(homeview.Props {
		Cookies: cookies,
	}))
}


