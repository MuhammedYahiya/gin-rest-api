package main

import (
	"github.com/gin-gonic/gin"
	"github.com/web-server-gin/routes"
)

func main() {
	r := gin.Default()
	routes.InitializeRoutes(r)

	r.Run(":8080")
}
