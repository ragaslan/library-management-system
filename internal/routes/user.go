package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ragaslan/library-management-system/internal/handlers"
)

func registerUser(r fiber.Router, h *handlers.Handler) {
	users := r.Group("/users")
}
