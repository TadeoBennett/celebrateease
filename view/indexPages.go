package view

import (
	// "errors"
	// "fmt"
	"html/template"
	// "log"
	"net/http"
)

type UserView struct{}

func (uc *UserView) RenderGuestIndexPage(w http.ResponseWriter) {
	//Display quotes using a template
	tmpl, err := template.ParseFiles("./views/guestIndex.tmpl")

    if err != nil {
        http.Error(w, "Failed to render template 18", http.StatusInternalServerError)
        return
    }

    err = tmpl.Execute(w, nil)
    if err != nil {
        http.Error(w, "Failed to show guest index page", http.StatusInternalServerError)
        return
    }
}
