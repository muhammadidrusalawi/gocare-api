package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadidrusalawi/gocare-api/internal/helper"
	"github.com/muhammadidrusalawi/gocare-api/internal/model"
	"github.com/muhammadidrusalawi/gocare-api/internal/repository"
	"github.com/muhammadidrusalawi/gocare-api/internal/request"
	"github.com/muhammadidrusalawi/gocare-api/provider/database"
)

func AdminGetAllCategories() []*model.Category {
	categoryRepo := repository.NewCategoryRepository(database.DB)

	categories, err := categoryRepo.FindAll()
	if err != nil {
		return []*model.Category{}
	}
	return categories
}

func AdminGetCategoryByID(ID string) (*model.Category, error) {
	categoryRepo := repository.NewCategoryRepository(database.DB)

	category, err := categoryRepo.FindByID(ID)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Category not found")
	}

	return category, nil
}

func AdminCreateCategory(req request.CreateCategoryRequest) (*model.Category, error) {
	categoryRepo := repository.NewCategoryRepository(database.DB)
	categoryExist, err := categoryRepo.FindByName(req.Name)
	if err == nil && categoryExist != nil {
		return nil, fiber.NewError(fiber.StatusConflict, "Category name already exists")
	}

	slug, err := helper.GenerateUniqueSlug(
		database.DB,
		&model.Category{},
		req.Name,
	)

	if err != nil {
		return nil, err
	}

	category := &model.Category{
		Name:        req.Name,
		Slug:        slug,
		Description: req.Description,
	}

	if err := categoryRepo.Create(category); err != nil {
		return nil, err
	}

	return category, nil
}

func AdminUpdateCategory(ID string, req request.UpdateCategoryRequest) (*model.Category, error) {
	categoryRepo := repository.NewCategoryRepository(database.DB)

	category, err := categoryRepo.FindByID(ID)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Category not found")
	}

	if req.Name != "" {
		categoryExist, err := categoryRepo.FindByName(req.Name)
		if err == nil && categoryExist != nil && categoryExist.ID != category.ID {
			return nil, fiber.NewError(fiber.StatusConflict, "Category name already exists")
		}

		slug, err := helper.GenerateUniqueSlug(
			database.DB,
			&model.Category{},
			req.Name,
		)
		if err != nil {
			return nil, err
		}

		category.Name = req.Name
		category.Slug = slug
	}

	if req.Description != nil {
		category.Description = req.Description
	}

	if err := categoryRepo.Update(category); err != nil {
		return nil, err
	}

	return category, nil
}

func AdminDeleteCategory(ID string) error {
	categoryRepo := repository.NewCategoryRepository(database.DB)

	_, err := categoryRepo.FindByID(ID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Category not found")
	}

	if err := categoryRepo.Delete(ID); err != nil {
		return err
	}

	return nil
}
