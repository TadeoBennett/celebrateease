package main

import (
	"net/http"
	"tadeobennett/celebrateease/controller"
	// "tadeobennett/celebrateease/model"
	"tadeobennett/celebrateease/view"
)

func (app *application) GuestHome(w http.ResponseWriter, r *http.Request) {
	//creates the controller with dependency injection
	userView := &view.UserView{}

	// Initialize UserController with injected dependencies
	uc := &controller.UserController{
		UserView: userView,
	}

	//calls the user controller to render the home page
	uc.RenderGuestHomePage(w, r)
}
