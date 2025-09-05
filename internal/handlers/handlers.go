package handlers

import (
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Handler struct {
	DB          *gorm.DB
	RDB         *redis.Client
	jwtSecret   string
	accessTTL   time.Duration
	refreshTTL  time.Duration
	JwtHandler  *JwtHandler
	AuthHandler *AuthHandler
}

type registerReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func New(
	db *gorm.DB,
	rdb *redis.Client,
	secret string,
	access time.Duration,
	refresh time.Duration,
) *Handler {
	JwtHandler := &JwtHandler{RDB: rdb, jwtSecret: secret, accessTTL: access, refreshTTL: refresh}
	AuthHandler := &AuthHandler{DB: db, JwtHandler: JwtHandler}

	return &Handler{
		DB:          db,
		RDB:         rdb,
		JwtHandler:  JwtHandler,
		AuthHandler: AuthHandler,
	}
}
