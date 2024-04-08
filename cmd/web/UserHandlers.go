package main

import (
	"net/http"
	"tadeobennett/celebrateease/controller"
	"tadeobennett/celebrateease/model/postgresql"

	// "tadeobennett/celebrateease/shared"
	// "tadeobennett/celebrateease/model"
	"tadeobennett/celebrateease/view"
)

func setupViewAndModelForUserController(app *Application) controller.UserController {
	// initialize view and model
	userView := &view.UserView{
		InfoLog:      app.InfoLog,
		ErrorLog:     app.ErrorLog,
		Session:      app.Session,
		ErrorHandler: app,
	}
	userModel := &postgresql.UserModel{
		DB: app.DB,
	}
	uc := &controller.UserController{
		UserView:  userView,
		UserModel: userModel,
	}
	return *uc
}

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	//setup the controller with the view and model
	uc := setupViewAndModelForUserController(app)
	//calls the user controller to render the home page
	uc.RenderHomePage(w, r)
}

func (app *Application) GuestHome(w http.ResponseWriter, r *http.Request) {
	//setup the controller with the view and model
	uc := setupViewAndModelForUserController(app)
	//calls the user controller to render the home page
	uc.RenderGuestHomePage(w, r)
}

func (app *Application) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	//setup the controller with the view and model
	uc := setupViewAndModelForUserController(app)
	uc.RenderAllUsers(w, r)
}
