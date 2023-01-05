package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// register middleware
	mux.Use(middleware.Recoverer)

	// mux.Use(app.enableCORS)

	// authentication routes - auth handler, refresh
	mux.Post("/auth", app.authenticate)
	mux.Post("/refresh-token", app.refresh)

	// test handler

	// protected routes
	mux.Route("/users", func(r chi.Router) {
		// user auth middleware

		r.Get("/", app.allUsers)
		r.Get("/{userID}", app.getUser)
		r.Delete("/{userID}", app.deleteUser)
		r.Put("/", app.insertUser)
		r.Patch("/", app.updateUser)
	})

	return mux
}
