package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"tadeobennett/celebrateease/initializers"

	"github.com/golangcollege/sessions"
	// "github.com/golangcollege/sessions"
)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Session  *sessions.Session
	DB       *sql.DB
}

func main() {
	//create flags
	addr := flag.String("addr", ":4000", "HTTP network address") //set a flag to use a custom port or port :4000 by default
	session := initializers.NewSession()
	flag.Parse()

	// initialize all necessities
	initializers.LoadEnvVariables()
	db, err := initializers.ConnectToDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	initializers.SyncDb()

	//create custome logs
	infoLog, errorLog := CreateCustomLog()

	app := &Application{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
		Session:  session,
		DB:       db,
	}

	// app.routes()

	// custom web http server
	srv := &http.Server{
		Addr:     *addr,
		Handler:  routes(app), //return the multiplexer
		ErrorLog: errorLog,    // initialize the standard error log with my own errorlog
	}

	infoLog.Printf("Starting server on port %s", *addr)
	srv.ListenAndServe()
	// err = srv.ListenAndServeTLS("../../tls/cert.pem", "../../tls/key.pem") //use the certifate values
	srv.ErrorLog.Fatal(err)
}
