package connectors

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

const DRIVER = "postgres"

func Create(credentials string) *sql.DB {
	db, err := sql.Open(DRIVER, credentials)
	errCheck(err)

	err = db.Ping()
	errCheck(err)

	return db
}

func errCheck(err error) {
	if (err != nil) {
		log.Fatal(err)
	}
}
