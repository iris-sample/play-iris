package models

import "time"

type Page struct {
	Id      uint `gorm:"primary_key"`
	Title   string `gorm:"type:varchar(40);"`
	Content string
	Created time.Time
}

func (p *Page) FormattedDate() string {
	return p.Created.Format("02/01/2006 15:04:05")
}