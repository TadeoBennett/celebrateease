package main

import (
	"net/http"
	"github.com/bmizerany/pat"
	// "github.com/justinas/alice"
)

func (app *application) routes() http.Handler{
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.GuestHome)) //show the home page with celebrants
	mux.Get("/guest", http.HandlerFunc(app.GuestHome)) //show the home page with celebrants

	//create a file server to serve out static content (css, js, sass, images, etc.)
	fileServer := http.FileServer(http.Dir("../../ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static/", fileServer))

	return mux
}
