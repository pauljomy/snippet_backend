package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", app.Home)
	mux.Get("/snippet/view/{id}", app.snippetView)
	mux.Get("/snippet/create", app.snippetCreate)
	mux.Post("/snippet/create", app.snippetCreatePost)
	return mux
}
