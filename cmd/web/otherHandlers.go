package main

import (
	"net/http"
	"tadeobennett/celebrateease/view"
)

func (app *Application) DisplayLoginPage(w http.ResponseWriter, r *http.Request) {
	//initialize a view using this application values
	view := &view.View{
		InfoLog:      app.InfoLog,
		ErrorLog:     app.ErrorLog,
		Session:      app.Session,
		ErrorHandler: app,
	}
	//run the view's function to display the login page
	view.ShowLoginPage(w, r)
}
