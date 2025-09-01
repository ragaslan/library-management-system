package handlers

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ragaslan/library-management-system/internal/models"
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
	return &Handler{DB: db, RDB: rdb, jwtSecret: secret, accessTTL: access, refreshTTL: refresh}
}

func (h *Handler) Register(c *fiber.Ctx) error {
	var req registerReq
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}
	if len(req.Password) < 6 || !strings.Contains(req.Email, "@") {
		return fiber.NewError(400, "invalid email or password")
	}
	hash, _ := hashPassword(req.Password)
	user := models.User{Email: req.Email, Password: hash}

	if err := h.DB.Create(&user).Error; err != nil {
		return fiber.NewError(400, "Register Error")
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"id":    user.ID,
			"email": user.Email,
		},
	})
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var in loginReq
	if err := c.BodyParser(&in); err != nil {
		return fiber.ErrBadRequest
	}

	var user models.User
	if err := h.DB.First(&user, "email = ?", in.Email).Error; err != nil {
		return fiber.ErrUnauthorized
	}

	if err := checkPassword(user.Password, in.Password); err != nil {
		return fiber.ErrUnauthorized
	}

	access, _, err := h.signAccess(user.ID, user.Role)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"token":   access,
	})
}
