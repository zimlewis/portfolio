package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zimlewis/portfolio/components"
)


func home(c *gin.Context) {
	c.HTML(http.StatusOK, "", components.Home())
}


