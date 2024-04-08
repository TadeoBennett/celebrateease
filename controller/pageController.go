package controller

import (
	"database/sql"
	"log"
	"net/http"

	// "tadeobennett/celebrateease/model"
	"tadeobennett/celebrateease/model/postgresql"
	"tadeobennett/celebrateease/shared"
	"tadeobennett/celebrateease/view"
	"tadeobennett/celebrateease/view/templates"
	// "github.com/golangcollege/sessions"
)

// PageController handles user-related operations
type PageController struct {
	PageModel *postgresql.PageModel
	PageView  *view.PageView
	DB        *sql.DB
	shared.ErrorHandler
	// errorLog  *log.Logger
	// infoLog   *log.Logger
	// session   *sessions.Session
}

func (uc *PageController) RenderAllPages(w http.ResponseWriter, r *http.Request) {
	//use the model here to get any necessary data from the database
	allPages, err := uc.PageModel.GetPagesFromDB()

	if err != nil {
		log.Panic(err.Error())
		uc.ServerError(w, err)
		return
	}

	//return errors or proceed
	data := &templates.PageTemplate{
		Pages: allPages,
	}

	// Display the created user in another template
	uc.PageView.RenderAllPagesOnPage(w, data)
}
