package database

import "database/sql"

type Storage struct {
	database *sql.DB
}

func CreateNewStorage(db *sql.DB) *Storage {
	return &Storage{
		database: db,
	}
}
