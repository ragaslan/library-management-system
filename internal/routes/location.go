package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ragaslan/library-management-system/internal/handlers"
)

func registerLocation(r fiber.Router, h *handlers.Handler) {
	locations := r.Group("/location")
	locations.Get("/", h.LocationHandler.GetAllLocations)
	locations.Post("/", h.LocationHandler.CreateLocation)
	locations.Delete("/:id", h.LocationHandler.DeleteLocationById)
}
