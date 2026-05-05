package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() error {
	var err error
	DB, err = sql.Open("mysql", os.Getenv("DB_DSN"))
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS urls (
			code       VARCHAR(6) PRIMARY KEY,
			url        TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP NULL DEFAULT NULL
		)
	`)
	if err != nil {
		return err
	}

	// Lägg till deleted_at om den saknas (ignorerar fel om kolumnen redan finns)
	DB.Exec("ALTER TABLE urls ADD COLUMN deleted_at TIMESTAMP NULL DEFAULT NULL")

	return nil
}
