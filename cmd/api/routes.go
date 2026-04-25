package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(app.logRequest)
	mux.Use(middleware.Recoverer)
	mux.Use(commonHeaders)
	mux.Use(enableCORS)

	mux.Get("/", app.Home)
	mux.Get("/snippet/view/{id}", app.snippetView)
	mux.Get("/snippet/create", app.snippetCreate)
	mux.Post("/snippet/create", app.snippetCreatePost)
	return mux
}
