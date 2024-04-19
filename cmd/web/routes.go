package main

import (
	// "log"

	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/bmizerany/pat"
	// "github.com/golangcollege/sessions"
	"github.com/justinas/alice"
)

func routes(app *Application) http.Handler {
	//create a variable to hold my middleware chain in order
	standardMiddleware := alice.New(
		app.securityHeadersMiddleware,
		app.logRequestMiddleware,
		app.recoverPanicMiddleware,
	)
	//loads and saves session data to and from the session cookie
	dynamicMiddleware := alice.New(app.Session.Enable)
	mux := pat.New()

	// Users Routes
	//	mux.Get("/users", dynamicMiddleware.ThenFunc(app.GetAllUsers)) // Get all users
	//mux.Get("/users/:userid", dynamicMiddleware.ThenFunc(app.GetUserByID)) // Get user by ID
	// mux.Post("/users", dynamicMiddleware.ThenFunc(app.CreateUser))         // Create a new user
	// mux.Put("/users/:userid", dynamicMiddleware.ThenFunc(app.UpdateUser))  // Update an existing user
	// mux.Del("/users/:userid", dynamicMiddleware.ThenFunc(app.DeleteUser))  // Delete an existing user

	// Events Routes (note: session variables can be used to check for userid and event id. Use to shorten routes	mux.Get("/users/:userid/events", dynamicMiddleware.ThenFunc(app.GetEventsForUser))                           // Get all events for a user
	//mux.Get("/events", dynamicMiddleware.ThenFunc(app.GetAllEvents)) // Get all events
	// mux.Get("/events/:eventid", dynamicMiddleware.ThenFunc(app.GetEventByID))                  // Get event by ID
	// mux.Get("/events/:eventid/users/:userid", dynamicMiddleware.ThenFunc(app.GetEventForUser)) //Get events for a user
	// mux.Get("/events/:eventid/users/celebrant/:celebrantid", dynamicMiddleware.ThenFunc(app.GetEventForCelebrant)) //Get events for a celebrant
	// mux.Post("/events", dynamicMiddleware.ThenFunc(app.CreateEvent))                           // Create a new event
	// mux.Put("/events/:eventid", dynamicMiddleware.ThenFunc(app.UpdateEventForUser))            // Update an existing event
	// mux.Del("/events/:eventid", dynamicMiddleware.ThenFunc(app.DeleteEvent))                   // Delete an existing event

	// Celebrants Routes
	//mux.Get("/celebrants", dynamicMiddleware.ThenFunc(app.GetAllCelebrants)) // Get all celebrants
	// mux.Get("/celebrant/:celebrantid", dynamicMiddleware.ThenFunc(app.GetCelebrantByID))                               // Get celbrant by ID
	// mux.Post("/celebrant", dynamicMiddleware.ThenFunc(app.CreateCelebrant))                                            // Create a new celebrant
	// mux.Put("/celebrants/:celebrantid", dynamicMiddleware.ThenFunc(app.UpdateCelebrantForUser))                        // Update an existing celebrant
	// mux.Del("/celebrants/:celebrantid", dynamicMiddleware.ThenFunc(app.DeleteCelebrant))                               // Delete an existing celebrant

	// Pages Routes
	//mux.Get("/pages", dynamicMiddleware.ThenFunc(app.GetAllPages)) // Get all pages
	// mux.Get("/page/:pageid", dynamicMiddleware.ThenFunc(app.GetPageByID))                               // Get page by ID
	// mux.Post("/page", dynamicMiddleware.ThenFunc(app.CreatePage))                                            // Create a new event
	// mux.Put("/pages/:pageid", dynamicMiddleware.ThenFunc(app.UpdatePageForUser))                        // Update an existing page
	// mux.Del("/page/:pageid", dynamicMiddleware.ThenFunc(app.DeletePage))                               // Delete an existing page
	//create a file server to serve out static content (css, js, sass, images, etc.)

	// Other Custom System Routes
	// mux.Get("/user/:userid/celebrants", dynamicMiddleware.ThenFunc(app.GetCelebrantsForLoggedInUser)) // Get the celebrants for the logged in user

	// for the loading of the default pages
	mux.Get("/login", dynamicMiddleware.ThenFunc(app.DisplayLoginPage))
	// mux.Post("/login", dynamicMiddleware.ThenFunc(app.LoginUser))
	mux.Get("/dashboard", dynamicMiddleware.ThenFunc(app.DisplayLoginPage))
	mux.Get("/celebrants", dynamicMiddleware.ThenFunc(app.DisplayLoginPage))
	mux.Get("/events", dynamicMiddleware.ThenFunc(app.DisplayLoginPage))
	mux.Get("/pages", dynamicMiddleware.ThenFunc(app.DisplayLoginPage))



	// mux.Get("/", dynamicMiddleware.ThenFunc(app.GuestHome))
	// mux.Get("/dashboard", dynamicMiddleware.ThenFunc(app.IndexHome))

	// Parse all template files using ParseGlob
	mytmpls, err := template.ParseGlob("../../views/**/*.tmpl")
	if err != nil {
		panic(err)
	}

	// Load additional templates specified in templateFiles
	templates := make(map[string]*template.Template)
	templateFiles := []string{
		"auth-confirm.tmpl",
		"auth-login.tmpl",
		"auth-register.tmpl",
		"auth-resetpw.tmpl",
		"calendar.tmpl",
		"chart-apexcharts.tmpl",
		"chart-chartjs.tmpl",
		"chart-inline.tmpl",
		"contacts-grid.tmpl",
		"contacts-list.tmpl",
		"contacts-new.tmpl",
		"dashboard-analytics.tmpl",
		"dashboard-saas.tmpl",
		"dashboard-sales.tmpl",
		"dashboard-system.tmpl",
		"datamaps.tmpl",
		"files-grid.tmpl",
		"files-list.tmpl",
		"form_advanced.tmpl",
		"form_elements.tmpl",
		"form_layouts.tmpl",
		"form_upload.tmpl",
		"form_validation.tmpl",
		"form_wizard.tmpl",
		"guestIndex.tmpl",
		"index-boxed.tmpl",
		"index-horizontal.tmpl",
		"index-vertical.tmpl",
		"index.tmpl",
		"page-404.tmpl",
		"page-500.tmpl",
		"page-blank.tmpl",
		"page-invoice.tmpl",
		"page-orders.tmpl",
		"page-timeline.tmpl",
		"profile-notification.tmpl",
		"profile-security.tmpl",
		"profile-settings.tmpl",
		"profile.tmpl",
		"support-center.tmpl",
		"support-faqs.tmpl",
		"support-ticket-detail.tmpl",
		"support-tickets.tmpl",
		"table_advanced.tmpl",
		"table_basic.tmpl",
		"table_datatables.tmpl",
		"ui-buttons.tmpl",
		"ui-color.tmpl",
		"ui-icons.tmpl",
		"ui-modals.tmpl",
		"ui-notification.tmpl",
		"ui-progress.tmpl",
		"ui-tabs-accordion.tmpl",
		"ui-typograpy.tmpl",
		"widgets.tmpl",
		"dashboard-login.tmpl",
	}

	for _, filename := range templateFiles {
		tmpl, err := template.ParseFiles("../../views/" + filename)
		if err != nil {
			panic(err)
		}
		templates[filename] = tmpl
	}

	// Register handlers for each template file
	for filename, templates := range templates {
		filename := filename // Capture loop variable
		mux.Get("/"+filename, dynamicMiddleware.ThenFunc(func(w http.ResponseWriter, r *http.Request) {
			templates = LoadViewsAndPartials(w, app)
			err := templates.ExecuteTemplate(w, filename, nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}))
	}

	// Optionally, print parsed files to log
	for _, t := range mytmpls.Templates() {
		fmt.Println(t.Name())
	}

	fileServer := http.FileServer(http.Dir("../../static/"))
	mux.Get("/static/", http.StripPrefix("/static/", fileServer))
	return standardMiddleware.Then(mux)
}

func LoadViewsAndPartials(w http.ResponseWriter, app *Application) *template.Template {
	//save .tmp files in partials
	partials, err := filepath.Glob("../../views/partials/*.tmpl")
	if err != nil {
		app.TemplateError(w, err.Error())
	}
	views, err := filepath.Glob("../../views/*.tmpl") // save .tmpl files in views
	if err != nil {
		app.TemplateError(w, err.Error())
	}
	allTemplates := append(partials, views...)

	templates, err := template.ParseFiles(allTemplates...)
	if err != nil {
		app.TemplateError(w, err.Error())
	}
	return templates
}
