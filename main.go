package main

import (
	"instagram/download/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Instagram Download Service Running",
		})
	})

	r.GET("/download", handlers.DownloadHandler)

	r.Run(":8080")
}
