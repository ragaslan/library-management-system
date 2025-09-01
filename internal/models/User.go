package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username        string `gorm:"uniqueIndex"`
	Name            string
	Email           string `gorm:"uniqueIndex;size:255"`
	Password        string `gorm:"size:255"`
	IsEmailVerified bool   `gorm:"default:false"`
	Role            string `gorm:"default:user"`
}
