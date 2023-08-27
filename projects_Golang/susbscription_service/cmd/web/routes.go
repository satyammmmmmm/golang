package main

import "net/http"

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(app.SessionLoad)

	//define appliction routes

	mux.Get("/", app.HomePage)
	mux.Get("/login", app.LoginPagePage)
	mux.Post("/login", app.PostLoginPage)
	mux.Get("/logout", app.LogOut)
	mux.Get("/register", app.RegisterPage)

	mux.Post("/register", app.PostRegisterPage)
	mux.Get("/activate", app.ActivateAccount)

	
	return mux
}
