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
	standardMiddleware := alice.New(
		app.securityHeadersMiddleware,
		app.logRequestMiddleware,
		app.recoverPanicMiddleware,
	)
	dynamicMiddleware := alice.New(app.Session.Enable, noSurf)
	mux := pat.New()

	mux.Get("/", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.home)) // home page with celebrants
	mux.Get("/signup", dynamicMiddleware.ThenFunc(app.signupForm)) // form to signup/insert a new user
	mux.Post("/signup", dynamicMiddleware.ThenFunc(app.signup)) // signup/insert a new user
	mux.Get("/login", dynamicMiddleware.ThenFunc(app.loginForm)) // form to login a user
	mux.Post("/login", dynamicMiddleware.ThenFunc(app.login)) // login a user
	mux.Post("/logout", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.logout)) // logout a user
	// mux.Get("/user/:id/settings", dynamicMiddleware.ThenFunc(app.settingsForm)) // gets a user's settings
	// mux.Post("/user/:id/settings", dynamicMiddleware.ThenFunc(app.updateSettings)) // add/updates a user's settings

	//---------------------------------------------------------------------

	//---------------------------------------------------------------------
	// GET        /events                                              getSystemEvents                     -- get all system events
	// GET        /user/:id/events                                     getAllEventsForUser                     -- get all system events created by a user
	// GET        /user/:id/celebrant/                                 getUserCelebrants                -- get all celebrants
	mux.Get("/userid/celebrant/ ", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.home)) //--get all celebrants
	// GET        /user/:id/celebrant/:id/event/                       getUserCelebrantEvents           -- get all events for a celebrant
	// GET        /user/:id/celebrant/:id/event/:id/page               getUserCelebrantEventPages       -- get all pages for an event, for a celebrant, for a user
	mux.Get("/userid/celebrantid/eventid/page", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.getUserCelebrantEventPages)) //-- get all pages for an event, for a celebrant, for a user

	// GET        /user/:id/page/                                      getUserPages                     -- get all created pages for a user

	//---------------------------------------------------------------------
	// GET        /user/:id/celebrant/:id                              getUserCelebrant                 -- get celebrant
	mux.Get("/userid/celebrantid", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.getUserCelebrant)) //-- get celebrant
	// GET        /user/:id/celebrant/add                              createCelebrantForm              -- form to add a new celebrant
	// POST       /user/:id/celebrant/add                              createCelebrant                  -- adds a new celebrant
	// GET        /user/:id/celebrant/:id/edit                         editCelebrantForm                -- form to edit a celebrant
	// POST       /user/:id/celebrant/:id/edit                         editCelebrant                    -- edits a celebrant
	// POST       /user/:id/celebrant/:id/delete                       deleteCelebrant                  -- deletes a celebrant
	
	//---------------------------------------------------------------------
	// GET         /user/:id/celebrant/:id/event/:id                   getUserCelebrantEvent           -- get an event
	mux.Get("/userid/celebrantid/eventid", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.getUserCelebrantEvent)) 	   //-- get an event
	// GET         /user/:id/celebrant/:id/event/add                   createCelebrantEventForm        -- form to add a new event for a celebrant
	// POST        /user/:id/celebrant/:id/event/add                   createCelebrantEvent            -- adds a new event for a celebrant
	// GET         /user/:id/celebrant/:id/event/:id/edit              editCelebrantEventForm          -- form to edit a new event for a celebrant
	// POST        /user/:id/celebrant/:id/event/:id/edit              editCelebrantEvent              -- edits a new event for a celebrant
	// POST        /user/:id/celebrant/:id/event/:id/delete            deleteCelebrantEvent            -- deletes an event

	//---------------------------------------------------------------------
	// GET         /user/:id/celebrant/:id/event/:id/page/:id          getUserCelebrantEventPage       -- get a page
	// GET         /user/:id/celebrant/:id/event/:id/page/add          createCelebrantEventPageForm    -- form to add a new page for an event
	// POST        /user/:id/celebrant/:id/event/:id/page/add          createCelebrantEventPage        -- adds a new page for an event
	// GET         /user/:id/celebrant/:id/event/:id/page/:id/edit     editCelebrantEventPageForm      -- form to edit a new event for a celebrant
	// POST        /user/:id/celebrant/:id/event/:id/page/:id/edit     editCelebrantEventPage          -- edits a page for an event
	// POST        /user/:id/celebrant/:id/event/:id/page/:id/delete   deleteCelebrantEventPage        -- deletes a page for an event

	//---------------------------------------------------------------------
	// //basic actions
	// POST        /user/:id/celebrant/:id/event/:id/page/:id/code/edit  -- edits the code
	// POST        /user/:id/celebrant/:id/event/:id/page/:id/status/0   -- toggle the enabled page button

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
