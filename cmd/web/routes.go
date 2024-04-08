package main

import (
	// "log"

	"net/http"

	"github.com/bmizerany/pat"
	// "github.com/golangcollege/sessions"
	"github.com/justinas/alice"
)

func (app *Application) routes() http.Handler {
	//create a variable to hold my middleware chain in order
	standardMiddleware := alice.New(
		app.logRequestMiddleware,
		app.securityHeadersMiddleware,
		app.recoverPanicMiddleware,
	)
	//loads and saves session data to and from the session cookie
	dynamicMiddleware := alice.New(app.Session.Enable)
	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.Home))

	// Users Routes
	mux.Get("/users", dynamicMiddleware.ThenFunc(app.GetAllUsers)) // Get all users
	// mux.Get("/users/:userid", dynamicMiddleware.ThenFunc(app.GetUserByID)) // Get user by ID
	// mux.Post("/users", dynamicMiddleware.ThenFunc(app.CreateUser))         // Create a new user
	// mux.Put("/users/:userid", dynamicMiddleware.ThenFunc(app.UpdateUser))  // Update an existing user
	// mux.Del("/users/:userid", dynamicMiddleware.ThenFunc(app.DeleteUser))  // Delete an existing user

	// Events Routes (note: session variables can be used to check for userid and event id. Use to shorten routes	mux.Get("/users/:userid/events", dynamicMiddleware.ThenFunc(app.GetEventsForUser))                           // Get all events for a user
	mux.Get("/events", dynamicMiddleware.ThenFunc(app.GetAllEvents)) // Get all events
	// mux.Get("/events/:eventid", dynamicMiddleware.ThenFunc(app.GetEventByID))                  // Get event by ID
	// mux.Get("/events/:eventid/users/:userid", dynamicMiddleware.ThenFunc(app.GetEventForUser)) //Get events for a user
	// mux.Get("/events/:eventid/users/celebrant/:celebrantid", dynamicMiddleware.ThenFunc(app.GetEventForCelebrant)) //Get events for a celebrant
	// mux.Post("/events", dynamicMiddleware.ThenFunc(app.CreateEvent))                           // Create a new event
	// mux.Put("/events/:eventid", dynamicMiddleware.ThenFunc(app.UpdateEventForUser))            // Update an existing event
	// mux.Del("/events/:eventid", dynamicMiddleware.ThenFunc(app.DeleteEvent))                   // Delete an existing event

	// Celebrants Routes
	mux.Get("/celebrants", dynamicMiddleware.ThenFunc(app.GetAllCelebrants)) // Get all celebrants
	// mux.Get("/celebrant/:celebrantid", dynamicMiddleware.ThenFunc(app.GetCelebrantByID))                               // Get celbrant by ID
	// mux.Post("/celebrant", dynamicMiddleware.ThenFunc(app.CreateCelebrant))                                            // Create a new celebrant
	// mux.Put("/celebrants/:celebrantid", dynamicMiddleware.ThenFunc(app.UpdateCelebrantForUser))                        // Update an existing celebrant
	// mux.Del("/celebrants/:celebrantid", dynamicMiddleware.ThenFunc(app.DeleteCelebrant))                               // Delete an existing celebrant

	// Pages Routes
	mux.Get("/pages", dynamicMiddleware.ThenFunc(app.GetAllPages)) // Get all pages
	// mux.Get("/page/:pageid", dynamicMiddleware.ThenFunc(app.GetPageByID))                               // Get page by ID
	// mux.Post("/page", dynamicMiddleware.ThenFunc(app.CreatePage))                                            // Create a new event
	// mux.Put("/pages/:pageid", dynamicMiddleware.ThenFunc(app.UpdatePageForUser))                        // Update an existing page
	// mux.Del("/page/:pageid", dynamicMiddleware.ThenFunc(app.DeletePage))                               // Delete an existing page
	//create a file server to serve out static content (css, js, sass, images, etc.)
	fileServer := http.FileServer(http.Dir("../../static/"))
	mux.Get("/static/", http.StripPrefix("/static/", fileServer))

	fileServer2 := http.FileServer(http.Dir("../../views/"))
	mux.Get("/views/", http.StripPrefix("/views/", fileServer2))


	// fileServer3 := http.FileServer(http.Dir("../../views/partials"))
	// mux.Get("/partials/", http.StripPrefix("/partials/", fileServer3))

	return standardMiddleware.Then(mux)
}
