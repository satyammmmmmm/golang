package app

import (
	"src/api/controlllers/polo"
	"src/api/controlllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	router.POST("/repository", repositories.CreateRepo)

	router.POST("/repositories", repositories.CreateRepo)

}
