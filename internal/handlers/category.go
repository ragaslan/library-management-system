package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ragaslan/library-management-system/internal/models"
	"gorm.io/gorm"
)

type CategoryHandler struct {
	DB *gorm.DB
}

func (h *CategoryHandler) GetAllCategories(c *fiber.Ctx) error {
	var categories []models.BookCategory
	if err := h.DB.Find(&categories).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Fetch all categories error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    categories,
	})
}

func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	var categoryReq models.BookCategory
	if err := c.BodyParser(&categoryReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Bad request. Invalid category body",
		})
	}

	if err := h.DB.Create(&categoryReq).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Creating new category is failed",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    categoryReq,
	})

}

func (h *CategoryHandler) DeleteCategoryById(c *fiber.Ctx) error {
	categoryId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Bad request. Invalid category id",
		})
	}
	// find category to delete
	var category models.BookCategory
	if err := h.DB.Find(&category, "id=?", categoryId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "There is no valid Book Category with this id",
		})
	}

	if err := h.DB.Delete(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Delete category error!",
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"success": true,
	})

}
