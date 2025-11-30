package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zimlewis/portfolio/utils"
	"github.com/zimlewis/portfolio/views/projectsview"
)


func projects(c *gin.Context) {	
	cookies := utils.GetCookies(c)
	fmt.Println(cookies)
	c.HTML(http.StatusOK, "", projectsview.ProjectViews(projectsview.Props {
		Cookies: cookies,
	}))
}
