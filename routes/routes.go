package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/web-server-gin/controllers"
)

func InitializeRoutes(r *gin.Engine) {
	albumRoutes := r.Group("/albums")
	{
		albumRoutes.GET("/", controllers.GetAllAlbums)
	}
}
