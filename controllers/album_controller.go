package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/web-server-gin/model"
)

func GetAllAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": model.Albums})
}

func PostAlbums(c *gin.Context) {
	var newAlbum model.Album

	if err := c.ShouldBind(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	model.Albums = append(model.Albums, newAlbum)
	c.JSON(http.StatusCreated, gin.H{"data": newAlbum})

}

func GetAlbumsById(c *gin.Context) {
	id := c.Param("id")
	for _, album := range model.Albums {
		if album.ID == id {
			c.JSON(http.StatusOK, gin.H{"data": album})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
