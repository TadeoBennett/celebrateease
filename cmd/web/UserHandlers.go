package main

import (
	"net/http"
)

func (app *Application) GuestHome(w http.ResponseWriter, r *http.Request) {
	//setup the controller with the view and model
	uc := setupViewAndModelForUserController(app)
	//calls the user controller to render the home page
	uc.RenderGuestHomePage(w, r)
}
func (app *Application) IndexHome(w http.ResponseWriter, r *http.Request) {
	//setup the controller with the view and model
	uc := setupViewAndModelForUserController(app)
	//calls the user controller to render the home page
	uc.RenderIndexDashboardPage(w, r)
}

// func (app *Application) GetAllUsers(w http.ResponseWriter, r *http.Request) {
// 	//setup the controller with the view and model
// 	uc := setupViewAndModelForUserController(app)
// 	uc.RenderAllUsers(w, r)
// }
// func (app *Application) GetUserByID(w http.ResponseWriter, r *http.Request) {
// 	//setup the controller with the view and model
// 	uc := setupViewAndModelForUserController(app)
// 	uc.RenderUserByID(w, r)
// }
