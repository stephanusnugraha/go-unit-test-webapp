package main

import (
	"encoding/gob"
	"flag"
	"github.com/alexedwards/scs/v2"
	"go-unit-test-webapp/pkg/data"
	"go-unit-test-webapp/pkg/repository"
	"go-unit-test-webapp/pkg/repository/dbrepo"
	"log"
	"net/http"
)

type application struct {
	Session *scs.SessionManager
	DSN     string
	DB      repository.DatabaseRepo
}

func main() {
	gob.Register(data.User{})

	// set up an app config
	app := application{}
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC+7 connect_timeout=5", "Postgres connection")
	flag.Parse()

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}

	// get a session manager
	app.Session = getSession()

	// print out a message
	log.Println("Starting server on port 8080...")

	// start the server
	err = http.ListenAndServe(":8080", app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
