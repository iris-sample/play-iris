package repositories

import (
	"database/sql"
	"github.com/serajam/play-iris/app/models"
	"log"
	"github.com/jinzhu/gorm"
)

type PageRepository struct {
	Gorm *gorm.DB
}

func (p PageRepository) GetPage(id int) *models.Page {
	page := models.Page{}
	p.Gorm.First(&page, id)

	return &page
}

func (p PageRepository) GetPagesWithLimit(limit int) []*models.Page {
	pages := make([]*models.Page, limit)
	p.Gorm.Limit(limit).Order("created DESC").Find(&pages)

	return pages
}

func checkErr(err error) {
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that ID.")
	case err != nil:
		log.Fatal(err)
	}
}
