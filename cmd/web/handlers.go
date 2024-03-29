package main

import (
	"net/http"
	"tadeobennett/celebrateease/controller"
	"tadeobennett/celebrateease/model/postgresql"

	// "tadeobennett/celebrateease/model"
	"tadeobennett/celebrateease/view"
)

func (app *Application) GuestHome(w http.ResponseWriter, r *http.Request) {
	//creates the controller with dependency injection
	userView := &view.UserView{}

	// Initialize UserController with injected dependencies
	uc := &controller.UserController{
		UserView: userView,
	}

	//calls the user controller to render the home page
	uc.RenderGuestHomePage(w, r)
}

func (app *Application) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// initialize view and model
	userView := &view.UserView{
		InfoLog:  app.InfoLog,
		ErrorLog: app.ErrorLog,
		
	}

	userModel := &postgresql.UserModel{
		DB: app.DB,
	}
	uc := &controller.UserController{
		UserView:  userView,
		UserModel: userModel,
	}

	uc.RenderAllUsers(w, r)
}
