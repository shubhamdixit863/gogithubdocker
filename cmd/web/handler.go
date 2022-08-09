package main

import "net/http"

func (app *App) RootHandler(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte("Apis Working  Fine"))

}
