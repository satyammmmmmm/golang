package main

import (
	"gin-ad/controller"
	"gin-ad/middleware"
	"gin-ad/service"
	"io"
	
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	server := gin.New()
	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth())

	server.GET("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())

	})
	server.POST("/videos", func(ctx *gin.Context) {
		err := VideoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "video input is valid ",
			})
		}

	})
	server.Run()

}
