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

// EventController handles user-related operations
type EventController struct {
	EventModel *postgresql.EventModel
	EventView  *view.EventView
	DB        *sql.DB
	shared.ErrorHandler
	// errorLog  *log.Logger
	// infoLog   *log.Logger
	// session   *sessions.Session
}
func (uc *EventController) RenderAllEvents(w http.ResponseWriter, r *http.Request) {
	//use the model here to get any necessary data from the database
	allEvents, err := uc.EventModel.GetEventsFromDB()

	if err != nil {
		log.Panic(err.Error())
		uc.ServerError(w, err)
		return
	}

	//return errors or proceed
	data := &templates.EventTemplate{
		Events: allEvents,
	}

	// Display the created user in another template
	uc.EventView.RenderAllEventsOnPage(w, data)
}
