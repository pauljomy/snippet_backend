package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
	DSN    string
}

func main() {

	var addr string
	var dsn string

	flag.StringVar(&addr, "addr", ":4000", "HTTP network port address")
	flag.StringVar(&dsn, "dsn", "host=localhost port=5434 user=web password=pass dbname=snippetbox timezone=UTC connect_timeout=5", "Postgres Database")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
		DSN:    dsn,
	}

	conn, err := app.connectDB()
	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	logger.Info("Starting server", "addr", addr)

	err = http.ListenAndServe(addr, app.routes())
	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}

}
