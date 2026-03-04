package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadidrusalawi/gocare-api/internal/helper"
	"github.com/muhammadidrusalawi/gocare-api/internal/service"
)

type CategoryRequest struct {
	Name string `json:"name" validate:"required,max=255"`
}

func AdminGetAllCategoriesHandler(c *fiber.Ctx) error {
	categories := service.AdminGetAllCategories()

	if len(categories) == 0 {
		return c.JSON([]*struct{}{})
	}

	var result []fiber.Map

	for _, cat := range categories {
		result = append(result, fiber.Map{
			"id":   cat.ID,
			"name": cat.Name,
			"slug": cat.Slug,
		})
	}

	return c.JSON(helper.ApiSuccess("Categories data successfully retrieved", result))
}

func AdminCreateCategoryHandler(c *fiber.Ctx) error {
	var req CategoryRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError("Invalid JSON"))
	}

	if err := helper.ValidateStruct(req); err != "" {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError(err))
	}

	category, err := service.AdminCreateCategory(req.Name)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(helper.ApiSuccess("Category created", fiber.Map{
		"id":   category.ID,
		"name": category.Name,
		"slug": category.Slug,
	}))
}

func AdminGetCategoryByIDHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	category, err := service.AdminGetCategoryByID(id)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("Category detail successfully retrieved", fiber.Map{
		"id":   category.ID,
		"name": category.Name,
		"slug": category.Slug,
	}))
}

func AdminUpdateCategoryHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	var req CategoryRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError("Invalid JSON"))
	}

	if err := helper.ValidateStruct(req); err != "" {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError(err))
	}

	category, err := service.AdminUpdateCategory(id, req.Name)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("Category updated successfully", fiber.Map{
		"id":   category.ID,
		"name": category.Name,
		"slug": category.Slug,
	}))
}

func AdminDeleteCategoryHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := service.AdminDeleteCategory(id); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("Category deleted successfully", nil))
}
