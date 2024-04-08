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

type PageView struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Session  *sessions.Session
	shared.ErrorHandler
}

func (uv *PageView) RenderAllPagesOnPage(w http.ResponseWriter, data *templates.PageTemplate) {

	tmpl, err := template.ParseFiles("../../views/pages.tmpl")

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
