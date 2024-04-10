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

type CelebrantView struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Session  *sessions.Session
	shared.ErrorHandler
}

func (uv *CelebrantView) RenderAllCelebrantsOnPage(w http.ResponseWriter, data *templates.CelebrantTemplate) {
	tmpl, err := template.ParseFiles("../../views/guestIndex.tmpl")
	if err != nil {
		uv.NotFound(w, err)
		return
	}
	
	//if there are no errors
	err = tmpl.Execute(w, data)
	
	if err != nil {
		uv.NotFound(w, err)
		return
	}
}
