package main

import (
	"net/http"
	"tadeobennett/celebrateease/templates"
	"text/template"

	"github.com/justinas/nosurf"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	//may need to use the signed in user id to get the celebrants
	ts, err := template.ParseFiles("../../views/pages/dashboard.tmpl",
		"../../views/partials/test_header.tmpl",
		"../../views/partials/test_nav.tmpl",
		"../../views/partials/test_aside.tmpl",
		"../../views/partials/scripts.tmpl",
	)

	if err != nil {
		app.ServerError(w, err)
		return
	}

	flash := app.Session.PopString(r, "flash")

	err = ts.Execute(w, &templates.UserTemplate{
		Flash:           flash,
		IsAuthenticated: app.IsAuthenticated(r),
		CSRFToken:       nosurf.Token(r),
	})
	if err != nil {
		app.ServerError(w, err)
	}
}

func (app *Application) getUserCelebrant(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("../../views/pages/celebrant.tmpl",
		"../../views/partials/test_header.tmpl",
		"../../views/partials/test_nav.tmpl",
		"../../views/partials/test_aside.tmpl",
		"../../views/partials/scripts.tmpl",
	)

	if err != nil {
		app.ServerError(w, err)
		return
	}

	flash := app.Session.PopString(r, "flash")

	err = ts.Execute(w, &templates.TemplateData{
		Flash:           flash,
		IsAuthenticated: app.IsAuthenticated(r),
		CSRFToken:       nosurf.Token(r),
	})
	if err != nil {
		app.ServerError(w, err)
	}
}

func (app *Application) getUserCelebrantEventPages(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("../../views/pages/eventPages.tmpl",
		"../../views/partials/test_header.tmpl",
		"../../views/partials/test_nav.tmpl",
		"../../views/partials/test_aside.tmpl",
		"../../views/partials/scripts.tmpl",
	)

	if err != nil {
		app.ServerError(w, err)
		return
	}	

	flash := app.Session.PopString(r, "flash")

	err = ts.Execute(w, &templates.TemplateData{
		Flash:           flash,
		IsAuthenticated: app.IsAuthenticated(r),
		CSRFToken:       nosurf.Token(r),
	})
	if err != nil {
		app.ServerError(w, err)
	}
}

func (app *Application) getUserCelebrantEvent(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("../../views/pages/event.tmpl",
		"../../views/partials/test_header.tmpl",
		"../../views/partials/test_nav.tmpl",
		"../../views/partials/test_aside.tmpl",
		"../../views/partials/scripts.tmpl",
	)

	if err != nil {
		app.ServerError(w, err)
		return
	}	

	flash := app.Session.PopString(r, "flash")

	err = ts.Execute(w, &templates.TemplateData{
		Flash:           flash,
		IsAuthenticated: app.IsAuthenticated(r),
		CSRFToken:       nosurf.Token(r),
	})
	if err != nil {
		app.ServerError(w, err)
	}
}
