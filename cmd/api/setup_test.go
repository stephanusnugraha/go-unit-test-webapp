package main

import (
	"go-unit-test-webapp/pkg/repository/dbrepo"
	"os"
	"testing"
)

var app application

func TestMain(m *testing.M) {
	app.DB = &dbrepo.TestDBRepo{}
	app.Domain = "example.com"
	app.JWTSecret = "FPPi978TZicURVnX683eF3lL9s30UsoUg2Ytp5Fj"
	os.Exit(m.Run())
}
