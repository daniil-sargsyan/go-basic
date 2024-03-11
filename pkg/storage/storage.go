package storage

import (
	"fmt"
	"github.com/daniil-sargsyan/go-basic/pkg/config"
	"log"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Storage interface {
	// add methods to storage and implement them
}

type storage struct {
	db *sqlx.DB
}

func NewStorage(cnf *config.DB) Storage {
	conn, err := sqlx.Connect("pgx", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cnf.DBHost,
		cnf.DBPort,
		cnf.DBUser,
		cnf.DBPass,
		cnf.DBName,
	))

	if err != nil {
		log.Fatalf("Error connecting to database: %s", err.Error())
	}

	return &storage{
		db: conn,
	}
}
