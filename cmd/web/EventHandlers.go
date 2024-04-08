package main

import (
	"net/http"
	"tadeobennett/celebrateease/controller"
	"tadeobennett/celebrateease/model/postgresql"

	// "tadeobennett/celebrateease/shared"
	// "tadeobennett/celebrateease/model"
	"tadeobennett/celebrateease/view"
)

func setupViewAndModelForEventController(app *Application) controller.EventController {
	// initialize view and model
	eventView := &view.EventView{
		InfoLog:      app.InfoLog,
		ErrorLog:     app.ErrorLog,
		Session:      app.Session,
		ErrorHandler: app,
	}
	eventModel := &postgresql.EventModel{
		DB: app.DB,
	}
	ec := &controller.EventController{
		EventView:  eventView,
		EventModel: eventModel,
	}
	return *ec
}

func (app *Application) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	//setup the controller with the view and model
	ec := setupViewAndModelForEventController(app)
	ec.RenderAllEvents(w, r)
}
