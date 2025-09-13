package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ISBN        string `gorm:"uniqueIndex"`
	Name        string
	AuthorID    *uint `gorm:"index"` // nullable olsun istiyorsan *uint kullan
	Author      Author
	PublishedAt time.Time
	PageCount   int
	CategoryID  *uint `gorm:"index"`
	LocationID  *uint `gorm:"index"`
}
