package handlers

import (
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Handler struct {
	DB         *gorm.DB
	RDB        *redis.Client
	jwtSecret  string
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func New(
	db *gorm.DB,
	rdb *redis.Client,
	secret string,
	access time.Duration,
	refresh time.Duration,
) *Handler {
	return &Handler{DB: db, RDB: rdb, jwtSecret: secret, accessTTL: access, refreshTTL: refresh}
}
