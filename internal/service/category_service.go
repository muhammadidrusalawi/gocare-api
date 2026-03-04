package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadidrusalawi/gocare-api/internal/helper"
	"github.com/muhammadidrusalawi/gocare-api/internal/model"
	"github.com/muhammadidrusalawi/gocare-api/internal/repository"
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

func AdminCreateCategory(name string) (*model.Category, error) {
	categoryRepo := repository.NewCategoryRepository(database.DB)
	categoryExist, err := categoryRepo.FindByName(name)
	if err == nil && categoryExist != nil {
		return nil, fiber.NewError(fiber.StatusConflict, "Category name already exists")
	}

	slug, err := helper.GenerateUniqueSlug(
		database.DB,
		&model.Category{},
		name,
	)

	if err != nil {
		return nil, err
	}

	category := &model.Category{
		Name: name,
		Slug: slug,
	}

	if err := categoryRepo.Create(category); err != nil {
		return nil, err
	}

	return category, nil
}

func AdminGetCategoryByID(ID string) (*model.Category, error) {
	categoryRepo := repository.NewCategoryRepository(database.DB)

	category, err := categoryRepo.FindByID(ID)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Category not found")
	}

	return category, nil
}

func AdminUpdateCategory(ID, name string) (*model.Category, error) {
	categoryRepo := repository.NewCategoryRepository(database.DB)

	category, err := categoryRepo.FindByID(ID)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Category not found")
	}

	categoryExist, err := categoryRepo.FindByName(name)
	if err == nil && categoryExist != nil {
		return nil, fiber.NewError(fiber.StatusConflict, "Category name already exists")
	}

	slug, err := helper.GenerateUniqueSlug(
		database.DB,
		&model.Category{},
		name,
	)

	if err != nil {
		return nil, err
	}

	category.Name = name
	category.Slug = slug

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
