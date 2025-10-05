package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/ragaslan/library-management-system/internal/db"
	"github.com/ragaslan/library-management-system/internal/handlers"
	"github.com/ragaslan/library-management-system/internal/routes"
	"github.com/redis/go-redis/v9"
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

	DB := db.CreateDB()
	rdb := redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_ADDR")})
	accessTTL := time.Duration(mustInt("ACCESS_TLL_MIN", 15)) * time.Minute
	refreshTTL := time.Duration(mustInt("REFRESH_TTL_DAYS", 30)) * 24 * time.Hour

	h := handlers.New(
		DB,
		rdb,
		os.Getenv("JWT_SECRET"),
		accessTTL,
		refreshTTL,
	)

	app := fiber.New()
	routes.Register(app, h)
	log.Println("Listening on : " + os.Getenv("PORT"))
	app.Listen(":" + os.Getenv("PORT"))
}
