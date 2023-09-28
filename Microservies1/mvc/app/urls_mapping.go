package app

import (
	"mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_Id", controllers.GetUser)
}
