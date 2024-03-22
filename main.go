package main

import (
	"flag"
	"log"
	"net/http"
	"tadeobennett/celebrateease/helpers"
	"tadeobennett/celebrateease/initializers"
)

type application struct {
	// celebrants *postgresql.CelebrantModel
	// eventpages *postgresql.EventPageModel
	// events     *postgresql.EventModel
	// users      *postgresql.UserModel
	errorLog *log.Logger
	infoLog  *log.Logger
	// session    *sessions.Session
}

func main() {
	//create flags
	addr := flag.String("addr", ":4000", "HTTP network address") //set a flag to use a custom port or port :4000 by default
	// dsn := flag.String("dsn", "postgres://"+dbname+":"+password+"@"+host+"/"+user+"?sslmode=disable", "PostgreSQL DSN (Data Source Name)")
	// secret := flag.String("secret", "8693b89c15217db6a4a90aa41cf0e6d5f31752aaf318b4e184f7c5a93a9a90c2", "Secret Key")
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
	infoLog, errorLog := helpers.CreateCustomLog()

	app := &application{
		// celebrants: &postgresql.CelebrantModel{
		// 	DB: db,
		// },
		// eventpages: &postgresql.EventPageModel{
		// 	DB: db,
		// },
		// events: &postgresql.EventModel{
		// 	DB: db,
		// },
		// users: &postgresql.UserModel{
		// 	DB: db,
		// },
		errorLog: errorLog,
		infoLog:  infoLog,
		// session:  session,
	}

	app.routes()

	// custom web http server
	srv := &http.Server{
		Addr: *addr,
		Handler:  app.routes(), //return the multiplexer
		ErrorLog: errorLog, // initialize the standard error log with my own errorlog
	}

	infoLog.Printf("Starting server on port %s", *addr)
	srv.ListenAndServe()
	// err = srv.ListenAndServeTLS("../../tls/cert.pem", "../../tls/key.pem") //use the certifate values
	// srv.ErrorLog.Fatal(err)
}
