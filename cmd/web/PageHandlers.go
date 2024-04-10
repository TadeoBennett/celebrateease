package main

import (
	"net/http"
	"tadeobennett/celebrateease/controller"
	"tadeobennett/celebrateease/model/postgresql"

	// "tadeobennett/celebrateease/shared"
	// "tadeobennett/celebrateease/model"
	"tadeobennett/celebrateease/view"
)

func setupViewAndModelForPageController(app *Application) controller.PageController {
	// initialize view and model
	pageView := &view.PageView{
		InfoLog:      app.InfoLog,
		ErrorLog:     app.ErrorLog,
		Session:      app.Session,
		ErrorHandler: app,
	}
	pageModel := &postgresql.PageModel{
		DB: app.DB,
	}
	pc := &controller.PageController{
		PageView:  pageView,
		PageModel: pageModel,
		Session: app.Session,
	}
	return *pc
}

func (app *Application) GetAllPages(w http.ResponseWriter, r *http.Request) {
	//setup the controller with the view and model
	pc := setupViewAndModelForPageController(app)
	pc.RenderAllPages(w, r)
}
