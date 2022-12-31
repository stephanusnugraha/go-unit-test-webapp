package main

import (
	"flag"
	"fmt"
	"go-unit-test-webapp/pkg/repository"
	"go-unit-test-webapp/pkg/repository/dbrepo"
	"log"
	"net/http"
)

const port = 8090

type application struct {
	DSN       string
	DB        repository.DatabaseRepo
	Domain    string
	JWTSecret string
}

func main() {
	var app application
	flag.StringVar(&app.Domain, "domain", "example.com", "Domain for application, e.g. company.com")
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC+7 connect_timeout=5", "Postgres connection")
	flag.StringVar(&app.JWTSecret, "jwt-secret", "FPPi978TZicURVnX683eF3lL9s30UsoUg2Ytp5Fj", "signing secret")
	flag.Parse()

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	app.DB = &dbrepo.PostgresDBRepo{
		DB: conn,
	}

	log.Printf("Starting api on port %d\n", port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
