package view

import (
	// "errors"
	// "fmt"
	"html/template"
	"log"

	// "log"
	"github.com/golangcollege/sessions"
	"net/http"
	"tadeobennett/celebrateease/view/templates"
)

type UserView struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Session  *sessions.Session
}

func (uc *UserView) RenderGuestIndexPage(w http.ResponseWriter) {
	//Display quotes using a template
	tmpl, err := template.ParseFiles("../../views/guestIndex.tmpl")

	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Failed to show guest index page", http.StatusInternalServerError)
		return
	}
}

func (uc *UserView) RenderAllUsersOnPage(w http.ResponseWriter, data *templates.UserTemplate) {

	tmpl, err := template.ParseFiles("../../views/users.tmpl")

	if err != nil {
		log.Println(err.Error())
		// app.serverError(w, err)
		return
	}

	//if there are no errors
	err = tmpl.Execute(w, data)

	if err != nil {
		log.Panicln(err.Error())
		// app.serverError(w, err)
		return
	}
}
