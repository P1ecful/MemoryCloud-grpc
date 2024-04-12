package database

import (
	"database/sql"
	"log"
)

type Repositiry struct {
	database *sql.DB
	log      *log.Logger
}

func CreateNewRepository(db *sql.DB, logger *log.Logger) *Repositiry {
	return &Repositiry{
		database: db,
		log:      logger,
	}
}

// !FIXME Create Connection to database
func (r *Repositiry) Connect() *sql.DB {
	var database *sql.DB
	return database
}
