package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	log.Log.Info("about to map url", "step:1", "status:pending")
	mapUrls()
	log.Log.Info("successfully map url", "step:1", "status:successful")

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
