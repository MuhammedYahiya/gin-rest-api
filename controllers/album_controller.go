package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/web-server-gin/model"
)

func GetAllAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": model.Albums})
}
