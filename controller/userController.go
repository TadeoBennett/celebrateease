package controller

import (
	"database/sql"
	"net/http"

	// "tadeobennett/celebrateease/model"

	"tadeobennett/celebrateease/model/postgresql"
	"tadeobennett/celebrateease/shared"
	"tadeobennett/celebrateease/view"

	"github.com/golangcollege/sessions"
	// "github.com/golangcollege/sessions"
)

// UserController handles user-related operations
type UserController struct {
	UserModel *postgresql.UserModel
	UserView  *view.UserView
	Session   *sessions.Session
	DB        *sql.DB
	shared.ErrorHandler
	// errorLog  *log.Logger
	// infoLog   *log.Logger
	// session   *sessions.Session
}

func (uc *UserController) RenderGuestHomePage(w http.ResponseWriter, r *http.Request) {
	//use the model here to get any necessary data from the database

	// Display the created user in another template
	uc.UserView.RenderGuestHomePage(w)
}




//from post requests
// func (uc *UserController) DisplayDashboardPage(w http.ResponseWriter, r *http.Request) {
// 	//use the model here to get any necessary data from the database

// }
