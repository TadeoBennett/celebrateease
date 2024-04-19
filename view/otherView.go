package view

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"tadeobennett/celebrateease/shared"

	"github.com/golangcollege/sessions"
)

type View struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Session  *sessions.Session
	shared.ErrorHandler
}

func (v *View) ShowLoginPage(w http.ResponseWriter, r *http.Request) {
	templates := LoadViewsAndPartialsForBasicView(w, v)
	err := templates.ExecuteTemplate(w, "dashboard-login.tmpl", nil) //execute base template
	if err != nil {
		v.ErrorHandler.NotFound(w, err)
	}
	for _, tmpl := range templates.Templates() { //show the parsed files in log
		fmt.Println(tmpl.Name())
	}
}

func LoadViewsAndPartialsForBasicView(w http.ResponseWriter, v *View) *template.Template {
	//save .tmp files in partials
	partials, err := filepath.Glob("../../views/partials/*.tmpl")
	if err != nil {
		v.ErrorHandler.TemplateError(w, err.Error())
	}
	components, err := filepath.Glob("../../views/components/*.tmpl")
	if err != nil {
		v.ErrorHandler.TemplateError(w, err.Error())
	}
	views, err := filepath.Glob("../../views/*.tmpl") // save .tmpl files in views
	if err != nil {
		v.ErrorHandler.TemplateError(w, err.Error())
	}
	allTemplates := append(partials, views...)
	allTemplates = append(allTemplates, components...)

	templates, err := template.ParseFiles(allTemplates...)
	if err != nil {
		v.ErrorHandler.TemplateError(w, err.Error())
	}
	return templates
}
