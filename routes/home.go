package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zimlewis/portfolio/views/homeview"
)


func home(c *gin.Context) {
	c.HTML(http.StatusOK, "", homeview.Home())
}


