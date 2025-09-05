package db

import (
	"log"
	"os"

	"github.com/ragaslan/library-management-system/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDB() *gorm.DB {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Book{},
		&models.Author{},
		&models.BookCategory{},
		&models.Location{},
	); err != nil {
		log.Fatal(err)
	}

	return db

}
