package handler

import (
	"context"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadidrusalawi/gocare-api/internal/helper"
	"github.com/muhammadidrusalawi/gocare-api/provider/storage"
)

const maxFileSize = 2 * 1024 * 1024

func ImageUploadHandler(c *fiber.Ctx) error {

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "file is required")
	}

	if fileHeader.Size > maxFileSize {
		return fiber.NewError(fiber.StatusBadRequest, "file size must be less than 2MB")
	}

	contentType := fileHeader.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		return fiber.NewError(fiber.StatusBadRequest, "only image files are allowed")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	cld, err := storage.NewCloudinary()
	if err != nil {
		return err
	}

	result, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
		Folder: "GoCare/Products/Uploads",
	})
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("Upload successfully", fiber.Map{
		"url":       result.SecureURL,
		"public_id": result.PublicID,
	}))
}

func ImageDeleteHandler(c *fiber.Ctx) error {

	type Request struct {
		PublicID string `json:"public_id"`
	}

	var req Request

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request")
	}

	if req.PublicID == "" {
		return fiber.NewError(fiber.StatusBadRequest, "public_id is required")
	}

	cld, err := storage.NewCloudinary()
	if err != nil {
		return err
	}

	_, err = cld.Upload.Destroy(context.Background(), uploader.DestroyParams{
		PublicID: req.PublicID,
	})

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(
		helper.ApiSuccess("Image deleted successfully", nil),
	)
}
