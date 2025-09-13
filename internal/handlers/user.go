package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ragaslan/library-management-system/internal/models"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	var users []models.User
	if err := h.DB.Find(&users).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, " Getting All User Error")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    users,
	})
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var reqBody models.User

	if err := c.BodyParser(&reqBody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, " Invalid User Data")
	}

	if err := h.DB.Create(&reqBody).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "  user create error")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    reqBody,
	})
}

func (h *UserHandler) GetUserById(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, " Invalid user id")
	}
	var user models.User

	if err := h.DB.First(&user, "id=?", userId).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, " Could not find any user ")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}

func (h *UserHandler) DeleteUserById(c *fiber.Ctx) error {
	userId, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, " Invalid user id ")
	}
	// find the user
	var user models.User
	if err := h.DB.First(&user, "id=?", userId).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, " Could not find any user ")
	}
	if err := h.DB.Delete(&user).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, " Delete user error ")
	}
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"success": true,
		"message": "User is successfully deleted.",
	})
}
