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

type EventView struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Session  *sessions.Session
	shared.ErrorHandler
}

func (uv *EventView) RenderAllEventsOnPage(w http.ResponseWriter, data *templates.EventTemplate) {

	tmpl, err := template.ParseFiles("../../views/events.tmpl")

	if err != nil {
		uv.ErrorHandler.ServerError(w, err)
		return
	}
	
	//if there are no errors
	err = tmpl.Execute(w, data)
	
	if err != nil {
		uv.ErrorHandler.ServerError(w, err)
		return
	}
}
