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

	"github.com/golangcollege/sessions"
	// "github.com/golangcollege/sessions"
)

// CelebrantController handles user-related operations
type CelebrantController struct {
	CelebrantModel *postgresql.CelebrantModel
	CelebrantView  *view.CelebrantView
	DB             *sql.DB
	Session        *sessions.Session
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

// func (uc *CelebrantController) RenderCelebrantsForLoggedInUser(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi((r.URL.Query().Get(":id"))) //gets the ID from the URL as an integer instead of a string
// 	celebrants, err := uc.CelebrantModel.GetCelebrantsForLoggedInUserFromDB(id)

// 	if err != nil {
// 		log.Panic(err.Error())
// 		uc.ServerError(w, err)
// 		return
// 	}

// 	//return errors or proceed
// 	data := &templates.CelebrantTemplate{
// 		Celebrants: celebrants,
// 	}

// 	// Display the created user in another template
// 	uc.CelebrantView.RenderCelebrantsForLoggedInUserOnPage(w, data)
// }
