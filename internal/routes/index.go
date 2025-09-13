package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ragaslan/library-management-system/internal/handlers"
)

func Register(app *fiber.App, h *handlers.Handler) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	registerAuth(v1, h)
	registerUser(v1, h)
	registerCategory(v1, h)
	registerLocation(v1, h)

}
