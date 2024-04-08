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

// UserController handles user-related operations
type UserController struct {
	UserModel *postgresql.UserModel
	UserView  *view.UserView
	DB        *sql.DB
	shared.ErrorHandler
	// errorLog  *log.Logger
	// infoLog   *log.Logger
	// session   *sessions.Session
}

func (uc *UserController) RenderHomePage(w http.ResponseWriter, r *http.Request) {
	//use the model here to get any necessary data from the database
	
	// Display the created user in another template
	uc.UserView.RenderIndexPage(w)
}
func (uc *UserController) RenderGuestHomePage(w http.ResponseWriter, r *http.Request) {
	//use the model here to get any necessary data from the database
	
	// Display the created user in another template
	uc.UserView.RenderGuestIndexPage(w)
}

func (uc *UserController) RenderAllUsers(w http.ResponseWriter, r *http.Request) {
	//use the model here to get any necessary data from the database
	allUsers, err := uc.UserModel.GetUsersFromDB()

	if err != nil {
		log.Panic(err.Error())
		uc.ServerError(w, err)
		return
	}

	//return errors or proceed
	data := &templates.UserTemplate{
		Users: allUsers,
	}

	// Display the created user in another template
	uc.UserView.RenderAllUsersOnPage(w, data)
}
