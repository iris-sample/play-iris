package connectors

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

const DRIVER = "postgres"

func Create(credentials string) *gorm.DB {
	db, err := gorm.Open(DRIVER, credentials)
	errCheck(err)

	err = db.DB().Ping()
	errCheck(err)

	if os.Getenv("DB_DEBUG") == "true" {
		db.LogMode(true)
	}

	return db
}

func errCheck(err error) {
	if (err != nil) {
		log.Fatal(err)
	}
}
