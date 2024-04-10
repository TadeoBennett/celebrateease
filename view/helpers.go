package view

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func LoadViewsAndPartials(w http.ResponseWriter, uv *UserView) *template.Template {
	//save .tmp files in partials
	partials, err := filepath.Glob("../../views/partials/*.tmpl")
	if err != nil {
		uv.ErrorHandler.TemplateError(w, err.Error())
	}
	components, err := filepath.Glob("../../views/components/*.tmpl")
	if err != nil {
		uv.ErrorHandler.TemplateError(w, err.Error())
	}
	views, err := filepath.Glob("../../views/*.tmpl") // save .tmpl files in views
	if err != nil {
		uv.ErrorHandler.TemplateError(w, err.Error())
	}
	allTemplates := append(partials, views...)
	allTemplates = append(allTemplates, components...)

	templates, err := template.ParseFiles(allTemplates...)
	if err != nil {
		uv.ErrorHandler.TemplateError(w, err.Error())
	}
	return templates
}