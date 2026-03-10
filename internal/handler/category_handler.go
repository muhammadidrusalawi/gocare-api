package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadidrusalawi/gocare-api/internal/helper"
	"github.com/muhammadidrusalawi/gocare-api/internal/repository"
	"github.com/muhammadidrusalawi/gocare-api/internal/request"
	"github.com/muhammadidrusalawi/gocare-api/internal/response"
	"github.com/muhammadidrusalawi/gocare-api/internal/service"
	"github.com/muhammadidrusalawi/gocare-api/provider/database"
)

func AdminGetAllCategoriesHandler(c *fiber.Ctx) error {
	categoryRepo := repository.NewCategoryRepository(database.DB)
	categories := service.AdminGetAllCategories()

	var result []response.CategoryResponse

	for _, cat := range categories {
		count, _ := categoryRepo.CountProducts(cat.ID)

		result = append(result, response.CategoryResponse{
			ID:           cat.ID,
			Name:         cat.Name,
			Slug:         cat.Slug,
			Description:  cat.Description,
			ProductCount: &count,
		})
	}

	return c.JSON(helper.ApiSuccess("Categories data successfully retrieved", result))
}

func AdminGetCategoryByIDHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	categoryRepo := repository.NewCategoryRepository(database.DB)
	category, err := service.AdminGetCategoryByID(id)

	if err != nil {
		return err
	}

	count, _ := categoryRepo.CountProducts(category.ID)
	createdAt := category.CreatedAt
	updatedAt := category.UpdatedAt

	res := response.CategoryResponse{
		ID:           category.ID,
		Name:         category.Name,
		Slug:         category.Slug,
		Description:  category.Description,
		ProductCount: &count,
		CreatedAt:    &createdAt,
		UpdatedAt:    &updatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("Category detail successfully retrieved", res))
}

func AdminCreateCategoryHandler(c *fiber.Ctx) error {
	var req request.CreateCategoryRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError("Invalid JSON"))
	}

	if err := helper.ValidateStruct(req); err != "" {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError(err))
	}

	category, err := service.AdminCreateCategory(req)

	if err != nil {
		return err
	}

	res := response.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Slug:        category.Slug,
		Description: category.Description,
		CreatedAt:   &category.CreatedAt,
		UpdatedAt:   &category.UpdatedAt,
	}

	return c.Status(fiber.StatusCreated).JSON(helper.ApiSuccess("Category created", res))
}

func AdminUpdateCategoryHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	var req request.UpdateCategoryRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError("Invalid JSON"))
	}

	if err := helper.ValidateStruct(req); err != "" {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError(err))
	}

	category, err := service.AdminUpdateCategory(id, req)

	if err != nil {
		return err
	}

	res := response.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Slug:        category.Slug,
		Description: category.Description,
		CreatedAt:   &category.CreatedAt,
		UpdatedAt:   &category.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("Category updated successfully", res))
}

func AdminDeleteCategoryHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := service.AdminDeleteCategory(id); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("Category deleted successfully", nil))
}
