package handlers

import (
	"instagram/download/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DownloadHandler(c *gin.Context) {
	postUrl := c.Query("url")
	if postUrl == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing url parameter",
		})
		return 
	}

	mediaUrl, err := services.GetInstagramMedia(postUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return 
	}

	c.JSON(http.StatusOK, gin.H{
		"media_url": mediaUrl,
	})
}
