package view

import (
	// "errors"
	// "fmt"
	"html/template"
	"log"

	// "log"
	"net/http"
	"tadeobennett/celebrateease/shared"
	"tadeobennett/celebrateease/view/templates"

	"github.com/golangcollege/sessions"
)

type UserView struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Session  *sessions.Session
	shared.ErrorHandler
}

func (uv *UserView) RenderGuestIndexPage(w http.ResponseWriter) {
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

func (uv *UserView) RenderAllUsersOnPage(w http.ResponseWriter, data *templates.UserTemplate) {

	tmpl, err := template.ParseFiles("../../views/uses.tmpl")

	if err != nil {
		log.Println(err.Error())
		uv.NotFound(w)
		return
	}
	
	//if there are no errors
	err = tmpl.Execute(w, data)
	
	if err != nil {
		log.Println(err.Error())
		uv.NotFound(w)
		return
	}
}
