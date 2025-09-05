package models

import "gorm.io/gorm"

type BookCategory struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
}
