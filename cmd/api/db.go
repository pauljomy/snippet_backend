package main

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func (app *application) openDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", app.DSN)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (app *application) connectDB() (*sql.DB, error) {
	db, err := app.openDB()
	if err != nil {
		return nil, err
	}
	app.logger.Info("Connected to postgres DB")

	return db, nil
}
