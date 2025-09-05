package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ragaslan/library-management-system/internal/models"
	"gorm.io/gorm"
)

type AuthHandler struct {
	DB         *gorm.DB
	JwtHandler *JwtHandler
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req registerReq
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}
	if len(req.Password) < 6 || !strings.Contains(req.Email, "@") {
		return fiber.NewError(400, "invalid email or password")
	}
	hash, _ := h.JwtHandler.hashPassword(req.Password)
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

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var in loginReq
	if err := c.BodyParser(&in); err != nil {
		return fiber.ErrBadRequest
	}

	var user models.User
	if err := h.DB.First(&user, "email = ?", in.Email).Error; err != nil {
		return fiber.ErrUnauthorized
	}

	if err := h.JwtHandler.checkPassword(user.Password, in.Password); err != nil {
		return fiber.ErrUnauthorized
	}

	access, _, err := h.JwtHandler.signAccess(user.ID, user.Role)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"token":   access,
	})
}
