package service

import (
	"context"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadidrusalawi/gocare-api/internal/helper"
	"github.com/muhammadidrusalawi/gocare-api/internal/model"
	"github.com/muhammadidrusalawi/gocare-api/internal/repository"
	"github.com/muhammadidrusalawi/gocare-api/internal/request"
	"github.com/muhammadidrusalawi/gocare-api/provider/database"
	"github.com/muhammadidrusalawi/gocare-api/provider/storage"
)

func AdminGetAllProducts() []*model.Product {
	productRepo := repository.NewProductRepository(database.DB)

	products, err := productRepo.FindAll()
	if err != nil {
		return []*model.Product{}
	}
	return products
}

func AdminGetProductByID(ID string) (*model.Product, error) {
	productRepo := repository.NewProductRepository(database.DB)

	product, err := productRepo.FindByID(ID)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Product not found")
	}

	return product, nil
}

func AdminCreateProduct(req request.CreateProductRequest) (*model.Product, error) {
	categoryRepo := repository.NewCategoryRepository(database.DB)
	productRepo := repository.NewProductRepository(database.DB)

	category, err := categoryRepo.FindByID(req.CategoryID)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Category not found")
	}

	slug, err := helper.GenerateUniqueSlug(
		database.DB,
		&model.Product{},
		req.Name,
	)

	if err != nil {
		return nil, err
	}

	product := &model.Product{
		Name:              req.Name,
		Slug:              slug,
		CategoryID:        category.ID,
		Thumbnail:         req.Thumbnail,
		ThumbnailPublicID: req.ThumbnailPublicID,
		Description:       req.Description,
		Stock:             req.Stock,
		Price:             req.Price,
		ExpirationDate:    &req.ExpirationDate,
	}

	if err := productRepo.Create(product); err != nil {
		return nil, err
	}

	return product, nil
}

func AdminUpdateProduct(ID string, req request.UpdateProductRequest) (*model.Product, error) {
	categoryRepo := repository.NewCategoryRepository(database.DB)
	productRepo := repository.NewProductRepository(database.DB)

	product, err := productRepo.FindByID(ID)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Product not found")
	}

	if req.CategoryID != "" {
		_, err := categoryRepo.FindByID(req.CategoryID)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusNotFound, "Category not found")
		}

		product.CategoryID = req.CategoryID
		product.Category = model.Category{}
	}

	if req.Name != "" {
		product.Name = req.Name
		slug, err := helper.GenerateUniqueSlug(database.DB, &model.Product{}, req.Name)
		if err != nil {
			return nil, err
		}
		product.Slug = slug
	}

	if req.Description != "" {
		product.Description = req.Description
	}

	if req.Thumbnail != "" {
		product.Thumbnail = req.Thumbnail
		product.ThumbnailPublicID = req.ThumbnailPublicID
	}

	if req.Stock != 0 {
		product.Stock = req.Stock
	}

	if req.Price != 0 {
		product.Price = req.Price
	}

	if !req.ExpirationDate.IsZero() {
		product.ExpirationDate = &req.ExpirationDate
	}

	if err := productRepo.Update(product); err != nil {
		return nil, err
	}

	product, err = productRepo.FindByID(product.ID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func AdminDeleteProduct(ID string, req request.DeleteProductRequest) error {
	productRepo := repository.NewProductRepository(database.DB)

	_, err := productRepo.FindByID(ID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Product not found")
	}

	if err := productRepo.Delete(ID); err != nil {
		return err
	}

	cld, err := storage.NewCloudinary()
	if err != nil {
		return err
	}

	_, err = cld.Upload.Destroy(context.Background(), uploader.DestroyParams{
		PublicID: req.ThumbnailPublicID,
	})

	if err != nil {
		return err
	}

	return nil
}
