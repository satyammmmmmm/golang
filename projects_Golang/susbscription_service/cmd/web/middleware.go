package main

import "net/http"

func (App *Config) SessionLoad(next http.Handler) http.Handler {
	return app.Session.LoadAndSave()
}
