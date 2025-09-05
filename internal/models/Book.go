package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ISBN        string `gorm:"uniqueIndex"`
	Name        string
	Author      Author
	PublishedAt time.Time
	PageCount   int
	Category    BookCategory
	Location    Location
}
