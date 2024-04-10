package main

import (
	"net/http"
	"tadeobennett/celebrateease/controller"
	"tadeobennett/celebrateease/model/postgresql"

	// "tadeobennett/celebrateease/shared"
	// "tadeobennett/celebrateease/model"
	"tadeobennett/celebrateease/view"
)

func setupViewAndModelForCelebrantController(app *Application) controller.CelebrantController {
	// initialize view and model
	celebrantView := &view.CelebrantView{
		InfoLog:      app.InfoLog,
		ErrorLog:     app.ErrorLog,
		Session:      app.Session,
		ErrorHandler: app,
	}
	celebrantModel := &postgresql.CelebrantModel{
		DB: app.DB,
	}
	cc := &controller.CelebrantController{
		CelebrantView:  celebrantView,
		CelebrantModel: celebrantModel,
		Session:        app.Session,
	}
	return *cc
}

func (app *Application) GetAllCelebrants(w http.ResponseWriter, r *http.Request) {
	//setup the controller with the view and model
	cc := setupViewAndModelForCelebrantController(app)
	cc.RenderAllCelebrants(w, r)
}
// func (app *Application) GetCelebrantsForLoggedInUser(w http.ResponseWriter, r *http.Request) {
// 	//setup the controller with the view and model
// 	cc := setupViewAndModelForCelebrantController(app)
// 	cc.RenderCelebrantsForLoggedInUser(w, r)
// }
