package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/pauljomy/snippet_backend/internals/models"
)

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {

	var addr string
	var dsn string

	flag.StringVar(&addr, "addr", ":4000", "HTTP network port address")
	flag.StringVar(&dsn, "dsn", "host=localhost port=5434 user=web password=pass dbname=snippetbox timezone=UTC connect_timeout=5", "Postgres Database")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	logger.Info("Connected to postgres DB")
	defer db.Close()

	app := &application{
		logger:   logger,
		snippets: &models.SnippetModel{DB: db},
	}

	logger.Info("Starting server", "addr", addr)

	err = http.ListenAndServe(addr, app.routes())
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

}
