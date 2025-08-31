package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string
	Name        string
	Email       string
	Password    string
	IsActivated bool
}
