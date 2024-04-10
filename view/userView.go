package view

import (
	// "errors"
	// "fmt"

	"fmt"
	"log"

	// "log"
	"net/http"
	"tadeobennett/celebrateease/shared"

	"github.com/golangcollege/sessions"
)

type UserView struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Session  *sessions.Session
	shared.ErrorHandler
}

func (uv *UserView) RenderGuestHomePage(w http.ResponseWriter) {

	templates := LoadViewsAndPartials(w, uv)
	err := templates.ExecuteTemplate(w, "guestIndex.tmpl", nil) //execute base template
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	for _, tmpl := range templates.Templates() { //show the parsed files in log
		fmt.Println(tmpl.Name())
	}
}
func (uv *UserView) RenderIndexDashboardPage(w http.ResponseWriter) {

	templates := LoadViewsAndPartials(w, uv)
	err := templates.ExecuteTemplate(w, "index.tmpl", nil) //execute base template
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	for _, tmpl := range templates.Templates() { //show the parsed files in log
		fmt.Println(tmpl.Name())
	}
}
