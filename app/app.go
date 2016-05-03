package app

import (
	"github.com/serajam/play-iris/app/connectors"
	"fmt"
	"database/sql"
	"log"
	"github.com/joho/godotenv"
	"os"
)

type Application struct {
	Db *sql.DB
}

func Initialize() Application {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	a := Application{
		connectors.Create(fmt.Sprintf("user=%s dbname=%s host=%s password=%s", os.Getenv("DB_USER"), os.Getenv("DB_DATABASE"), os.Getenv("DB_HOST"), os.Getenv("DB_PASSWORD"))),
	}

	return a
}