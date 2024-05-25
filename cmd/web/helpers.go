package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug" //able to see stacktrace when there are errors
	"strconv"
)

func CreateCustomLog() (infoLog *log.Logger, errorLog *log.Logger) {
	//Create a logger
	//logs anything that not an error.
	myInfoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	//logs an error
	myErrorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return myInfoLog, myErrorLog
}

// ---------------------- CUSTOM LOGS -----------------------//

func (app *Application) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Output(2, trace) //2 means that if there's an error we want the linenumber and file to be the caller not the callee

	log.Printf("error found %v", trace)
	//deal with the error status
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) ClientError(w http.ResponseWriter, status int) {
	trace := fmt.Sprintf("Client error with status %s", strconv.Itoa(status))
	app.ErrorLog.Output(2, trace)
	http.Error(w, http.StatusText(status), status)
}

func (app *Application) NotFound(w http.ResponseWriter, err error) {
	//can redirect to the error page
	app.ErrorLog.Println(err.Error())
	app.ClientError(w, http.StatusNotFound)
}

func (app *Application) TemplateError(w http.ResponseWriter, err string) {
	trace := fmt.Sprintf("%s\n%s", err, debug.Stack())
	app.ErrorLog.Output(2, trace) //2 means that if there's an error we want the linenumber and file to be the caller not the callee
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) IsAuthenticated(r *http.Request) bool {
	return app.Session.Exists(r, "authenticatedUserId")
}
