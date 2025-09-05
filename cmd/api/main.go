package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/ragaslan/library-management-system/internal/handlers"
	"github.com/ragaslan/library-management-system/internal/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func mustInt(env string, def int) int {
	v := os.Getenv(env)
	if v == "" {
		return def
	}
	i, _ := strconv.Atoi(v)
	return i
}

func main() {
	_ = godotenv.Load()

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

	rdb := redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_ADDR")})
	accessTTL := time.Duration(mustInt("ACCESS_TLL_MIN", 15)) * time.Minute
	refreshTTL := time.Duration(mustInt("REFRESH_TTL_DAYS", 30)) * 24 * time.Hour

	h := handlers.New(db, rdb, os.Getenv("JWT_SECRET"), accessTTL, refreshTTL)

	app := fiber.New()

	app.Post("/auth/register", h.Register)
	app.Post("/auth/login", h.Login)

	log.Println("Listening on : " + os.Getenv("APP_PORT"))
	app.Listen(os.Getenv("APP_PORT"))
}
