package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zimlewis/portfolio/routes"
)

func main() {
	server := gin.Default()
	server.Static("/static", "./public/static")

	routes.InitRoutes(server)

	server.Run()
}
