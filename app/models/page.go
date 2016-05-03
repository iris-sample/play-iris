package models

import "time"

type Page struct {
	Id      int
	Title   string
	Content string
	Created time.Time
}
