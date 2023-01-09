package main

import (
	"go-unit-test-webapp/pkg/repository/dbrepo"
	"os"
	"testing"
)

var app application
var expiredToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiYXVkIjoiZXhhbXBsZS5jb20iLCJleHAiOjE2NzI4OTU0MzksImlzcyI6ImV4YW1wbGUuY29tIiwibmFtZSI6IkpvaG4gRG9lIiwic3ViIjoiMSJ9.XMvuppbXYMSyBJoll-001L5CPrU1W_Tx0hXcT0MnmO8"

func TestMain(m *testing.M) {
	app.DB = &dbrepo.TestDBRepo{}
	app.Domain = "example.com"
	app.JWTSecret = "FPPi978TZicURVnX683eF3lL9s30UsoUg2Ytp5Fj"
	os.Exit(m.Run())
}
