package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadidrusalawi/gocare-api/internal/helper"
)

func AdminDashboardHandler(c *fiber.Ctx) error {
	return c.JSON(helper.ApiSuccess("Welcome To Dashboard", nil))
}
