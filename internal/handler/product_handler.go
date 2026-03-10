package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadidrusalawi/gocare-api/internal/helper"
	"github.com/muhammadidrusalawi/gocare-api/internal/request"
	"github.com/muhammadidrusalawi/gocare-api/internal/response"
	"github.com/muhammadidrusalawi/gocare-api/internal/service"
)

func AdminGetAllProductsHandler(c *fiber.Ctx) error {
	products := service.AdminGetAllProducts()

	var result []response.ProductResponse

	for _, prod := range products {
		res := response.ProductResponse{
			ID:                prod.ID,
			Name:              prod.Name,
			Slug:              prod.Slug,
			Description:       prod.Description,
			Thumbnail:         prod.Thumbnail,
			ThumbnailPublicID: prod.ThumbnailPublicID,
			Stock:             prod.Stock,
			Price:             prod.Price,
			ExpirationDate:    prod.ExpirationDate,
		}

		if prod.Category.ID != "" {
			res.Category = &response.CategoryResponse{
				ID:   prod.Category.ID,
				Name: prod.Category.Name,
				Slug: prod.Category.Slug,
			}
		}

		result = append(result, res)
	}

	return c.Status(fiber.StatusOK).JSON(
		helper.ApiSuccess("Products data successfully retrieved", result),
	)
}

func AdminGetProductByIDHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	product, err := service.AdminGetProductByID(id)

	if err != nil {
		return err
	}

	res := response.ProductResponse{
		ID:             product.ID,
		Name:           product.Name,
		Slug:           product.Slug,
		Description:    product.Description,
		Thumbnail:      product.Thumbnail,
		Stock:          product.Stock,
		Price:          product.Price,
		ExpirationDate: product.ExpirationDate,
		CreatedAt:      &product.CreatedAt,
		UpdatedAt:      &product.UpdatedAt,
	}

	if product.Category.ID != "" {
		res.Category = &response.CategoryResponse{
			ID:   product.Category.ID,
			Name: product.Category.Name,
			Slug: product.Category.Slug,
		}
	}

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("Product detail successfully retrieved", res))
}

func AdminCreateProductHandler(c *fiber.Ctx) error {
	var req request.CreateProductRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(helper.ApiError("Invalid JSON"))
	}

	if err := helper.ValidateStruct(req); err != "" {
		return c.Status(fiber.StatusBadRequest).
			JSON(helper.ApiError(err))
	}

	product, err := service.AdminCreateProduct(req)
	if err != nil {
		return err
	}

	res := response.ProductResponse{
		ID:             product.ID,
		Name:           product.Name,
		Slug:           product.Slug,
		Description:    product.Description,
		Thumbnail:      product.Thumbnail,
		Stock:          product.Stock,
		Price:          product.Price,
		ExpirationDate: product.ExpirationDate,
		CreatedAt:      &product.CreatedAt,
	}

	if product.Category.ID != "" {
		res.Category = &response.CategoryResponse{
			ID:   product.Category.ID,
			Name: product.Category.Name,
			Slug: product.Category.Slug,
		}
	}

	return c.Status(fiber.StatusCreated).
		JSON(helper.ApiSuccess("Product created", res))
}

func AdminUpdateProductHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	var req request.UpdateProductRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(helper.ApiError("Invalid JSON"))
	}

	if err := helper.ValidateStruct(req); err != "" {
		return c.Status(fiber.StatusBadRequest).
			JSON(helper.ApiError(err))
	}

	product, err := service.AdminUpdateProduct(id, req)
	if err != nil {
		return err
	}

	res := response.ProductResponse{
		ID:             product.ID,
		Name:           product.Name,
		Slug:           product.Slug,
		Description:    product.Description,
		Thumbnail:      product.Thumbnail,
		Stock:          product.Stock,
		Price:          product.Price,
		ExpirationDate: product.ExpirationDate,
		CreatedAt:      &product.CreatedAt,
		UpdatedAt:      &product.UpdatedAt,
	}

	if product.Category.ID != "" {
		res.Category = &response.CategoryResponse{
			ID:   product.Category.ID,
			Name: product.Category.Name,
			Slug: product.Category.Slug,
		}
	}

	return c.Status(fiber.StatusOK).
		JSON(helper.ApiSuccess("Product updated", res))
}

func AdminDeleteProductHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	var req request.DeleteProductRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(helper.ApiError("Invalid JSON"))
	}

	if err := helper.ValidateStruct(req); err != "" {
		return c.Status(fiber.StatusBadRequest).
			JSON(helper.ApiError(err))
	}

	if err := service.AdminDeleteProduct(id, req); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("Product deleted successfully", nil))
}
