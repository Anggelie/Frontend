package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func setupDatabase(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	// Crear tabla si no existe
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT NOT NULL UNIQUE,
            password_hash TEXT NOT NULL
        )
    `)
	if err != nil {
		return nil, err
	}

	log.Println(" Conectado a la base de datos y tabla lista.")
	return db, nil
}
