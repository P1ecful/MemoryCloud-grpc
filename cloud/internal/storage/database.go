package database

import (
	"database/sql"
	"fmt"
	"log"
	"memorycloud/cloud/internal/config"
)

type Repository struct {
	config *config.PostgresConnection
	log    *log.Logger
}

func CreateNewRepository(cfg *config.PostgresConnection, logger *log.Logger) *Repository {
	return &Repository{
		config: cfg,
		log:    logger,
	}
}

func (r *Repository) Connect() *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		r.config.Host, r.config.Port, r.config.Username, r.config.Password, r.config.Database))

	if err != nil {
		r.log.Fatal(err)
	}

	return db
}
