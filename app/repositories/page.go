package repositories

import (
	"database/sql"
	"github.com/serajam/play-iris/app/models"
	"log"
)

type PageRepository struct {
	Db *sql.DB
}

func (p PageRepository) GetPage(id int) *models.Page {
	page := models.Page{}
	row := p.Db.QueryRow("SELECT id, title, content FROM pages WHERE id = $1", id)
	err := row.Scan(&page.Id, &page.Title, &page.Content)

	checkErr(err)

	return &page
}

func checkErr(err error) {
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that ID.")
	case err != nil:
		log.Fatal(err)
	}
}
