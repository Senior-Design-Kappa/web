package backend

import (
	"database/sql"
	"log"

	"github.com/Senior-Design-Kappa/web/config"
	_ "github.com/mattn/go-sqlite3"
)

type Backend interface {
}

type backend struct {
	Config config.Config
	db     *sql.DB
}

func NewBackend(conf config.Config) (Backend, error) {
	db, err := sql.Open("sqlite3", "./db/data.db")
	if err != nil {
		log.Printf("error: could not open db connection (%+v)\n", err)
	}
	b := &backend{
		Config: conf,
		db:     db,
	}
	return b, nil
}
