package app

import (
	"book/controllers/ping"
	"book/controllers/users"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.GET("/internal/users/search", users.Search)
	router.PUT("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.POST("users/login", users.Login)

}
