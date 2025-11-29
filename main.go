package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/zimlewis/portfolio/routes"
)

func main() {
	// Load .env only on test environment, otherwise use environment variables
	if gin.Mode() == "debug" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	server := gin.Default()

	routes.InitRoutes(server)

	server.Run()
}
