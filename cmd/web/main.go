package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"
	"tadeobennett/celebrateease/initializers"
	"tadeobennett/celebrateease/pkg/models/postgresql"
	"time"

	"github.com/golangcollege/sessions"
	// "github.com/golangcollege/sessions"
)

type Application struct {
	events     *postgresql.EventModel
	celebrants *postgresql.CelebrantModel
	pages      *postgresql.PageModel
	users      *postgresql.UserModel
	ErrorLog   *log.Logger
	InfoLog    *log.Logger
	Session    *sessions.Session
}

func main() {
	//create flags
	addr := flag.String("addr", ":4000", "HTTP network address") //set a flag to use a custom port or port :4000 by default
	secret := flag.String("secret", "8693b89c15217db6a4a90aa41cf0e6d5f31752aaf318b4e184f7c5a93a9a90c2", "Secret Key")
	session := initializers.NewSession(*secret)
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

	//setup the TLS Config
	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	app := &Application{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
		Session:  session,
		events: &postgresql.EventModel{
			DB: db,
		},
		celebrants: &postgresql.CelebrantModel{
			DB: db,
		},
		pages: &postgresql.PageModel{
			DB: db,
		},
		users: &postgresql.UserModel{
			DB: db,
		},
	}

	// custom web http server
	srv := &http.Server{
		Addr:         *addr,
		Handler:      routes(app),
		ErrorLog:     errorLog,
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second, //max time taken to read a request
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Starting server on port %s", *addr)
	err = srv.ListenAndServeTLS("../../tls/cert.pem", "../../tls/key.pem") //use the certifate values
	srv.ErrorLog.Fatal(err)
}
