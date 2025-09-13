package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ragaslan/library-management-system/internal/handlers"
)

func registerCategory(r fiber.Router, h *handlers.Handler) {
	categories := r.Group("/category")
	categories.Get("/", h.CategoryHandler.GetAllCategories)
	categories.Post("/", h.CategoryHandler.CreateCategory)
	categories.Delete("/:id", h.CategoryHandler.DeleteCategoryById)
}
