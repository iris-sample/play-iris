package app

import (
	"github.com/serajam/play-iris/app/connectors"
	"fmt"
	"log"
	"github.com/joho/godotenv"
	"os"
	"github.com/jinzhu/gorm"
	"github.com/serajam/play-iris/app/routes"
	"github.com/serajam/play-iris/app/repositories"
)

type Application struct {
	Gorm *gorm.DB
}

func Initialize() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	a := Application{
		connectors.Create(fmt.Sprintf("user=%s dbname=%s host=%s password=%s", os.Getenv("DB_USER"), os.Getenv("DB_DATABASE"), os.Getenv("DB_HOST"), os.Getenv("DB_PASSWORD"))),
	}

	r := repositories.PageRepository{a.Gorm}
	routes.PagesRoutes(r)
}
