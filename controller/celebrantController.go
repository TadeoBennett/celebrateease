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

// CelebrantController handles user-related operations
type CelebrantController struct {
	CelebrantModel *postgresql.CelebrantModel
	CelebrantView  *view.CelebrantView
	DB        *sql.DB
	shared.ErrorHandler
	// errorLog  *log.Logger
	// infoLog   *log.Logger
	// session   *sessions.Session
}

func (uc *CelebrantController) RenderAllCelebrants(w http.ResponseWriter, r *http.Request) {
	//use the model here to get any necessary data from the database
	allCelebrants, err := uc.CelebrantModel.GetCelebrantsFromDB()

	if err != nil {
		log.Panic(err.Error())
		uc.ServerError(w, err)
		return
	}

	//return errors or proceed
	data := &templates.CelebrantTemplate{
		Celebrants: allCelebrants,
	}

	// Display the created user in another template
	uc.CelebrantView.RenderAllCelebrantsOnPage(w, data)
}
