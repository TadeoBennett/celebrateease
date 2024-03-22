package controller

import (
	"net/http"
	// "tadeobennett/celebrateease/model"
	"tadeobennett/celebrateease/view"
	"tadeobennett/celebrateease/model/postgresql"
)

// UserController handles user-related operations
type UserController struct {
	UserModel *postgresql.UserModel
	UserView  *view.UserView
}

// NewUser creates a new user
func (uc *UserController) RenderGuestHomePage(w http.ResponseWriter, r *http.Request) {
	//use the model here to get any necessary data from the database

	// Display the created user in another template
	uc.UserView.RenderGuestIndexPage(w)
}
