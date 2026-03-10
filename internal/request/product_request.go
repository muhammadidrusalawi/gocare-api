package request

import "time"

type CreateProductRequest struct {
	Name              string    `json:"name" validate:"required,max=255"`
	CategoryID        string    `json:"category_id" validate:"required,max=255"`
	Description       string    `json:"description" validate:"required,max=1000"`
	Thumbnail         string    `json:"thumbnail" validate:"required"`
	ThumbnailPublicID string    `json:"thumbnail_public_id" validate:"required"`
	Stock             uint      `json:"stock" validate:"required"`
	Price             uint      `json:"price" validate:"required"`
	ExpirationDate    time.Time `json:"expiration_date"`
}

type UpdateProductRequest struct {
	Name              string    `json:"name" validate:"omitempty,max=255"`
	CategoryID        string    `json:"category_id" validate:"omitempty,max=255"`
	Description       string    `json:"description" validate:"omitempty,max=1000"`
	Thumbnail         string    `json:"thumbnail" validate:"omitempty"`
	ThumbnailPublicID string    `json:"thumbnail_public_id" validate:"omitempty"`
	Stock             uint      `json:"stock" validate:"omitempty"`
	Price             uint      `json:"price" validate:"omitempty"`
	ExpirationDate    time.Time `json:"expiration_date" validate:"omitempty"`
}

type DeleteProductRequest struct {
	ThumbnailPublicID string `json:"thumbnail_public_id" validate:"required"`
}
