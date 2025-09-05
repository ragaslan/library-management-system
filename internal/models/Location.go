package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	Floor int
	Shelf int
}
