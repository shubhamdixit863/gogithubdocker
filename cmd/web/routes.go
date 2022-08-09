package main

import (
	"github.com/gorilla/mux"
)

func (app *App) registerRoutes() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/", app.RootHandler).Methods("GET")

	return r

}
