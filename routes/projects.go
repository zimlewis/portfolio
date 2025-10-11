package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zimlewis/portfolio/views/projectsview"
)


func projects(c *gin.Context) {	
	c.HTML(http.StatusOK, "", projectsview.ProjectViews())
}
