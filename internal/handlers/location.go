package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ragaslan/library-management-system/internal/models"
	"gorm.io/gorm"
)

type LocationHandler struct {
	DB *gorm.DB
}

func (h *LocationHandler) GetAllLocations(c *fiber.Ctx) error {
	var locations []models.Location
	if err := h.DB.Find(&locations).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Fetch all location error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    locations,
	})
}

func (h *LocationHandler) CreateLocation(c *fiber.Ctx) error {
	var locationReq models.Location
	if err := c.BodyParser(&locationReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Bad request. Invalid Location body",
		})
	}

	if err := h.DB.Create(&locationReq).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Creating new location is failed",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    locationReq,
	})

}

func (h *LocationHandler) DeleteLocationById(c *fiber.Ctx) error {
	locationId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Bad request. Invalid location id",
		})
	}
	// find category to delete
	var location models.Location
	if err := h.DB.Find(&location, "id=?", locationId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "There is no valid Location with this id",
		})
	}

	if err := h.DB.Delete(&location).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Delete Location error!",
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"success": true,
	})

}
