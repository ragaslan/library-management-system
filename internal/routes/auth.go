package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ragaslan/library-management-system/internal/handlers"
)

func registerAuth(r fiber.Router, h *handlers.Handler) {
	auth := r.Group("/auth")
	auth.Post("/auth/register", h.AuthHandler.Register)
	auth.Post("/auth/login", h.AuthHandler.Login)
}
