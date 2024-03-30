package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug" //able to see stacktrace when there are errors	

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

	//deal with the error status
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) ClientError(w http.ResponseWriter, status int) {
	// log.Printf("reached an int is %v", status)
	http.Error(w, http.StatusText(status), status)
}

func (app *Application) NotFound(w http.ResponseWriter) {
	//can redirect to the error page
	app.ClientError(w, http.StatusNotFound)
}
