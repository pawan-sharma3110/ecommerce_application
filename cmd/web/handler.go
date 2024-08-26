package main

import (
	"net/http"
)

func (app *application) VirtualTeminal(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("Hit the handler")
}
