package response

import "time"

type ProductResponse struct {
	ID                string            `json:"id"`
	Name              string            `json:"name"`
	Slug              string            `json:"slug"`
	Description       string            `json:"description,omitempty"`
	Thumbnail         string            `json:"thumbnail"`
	ThumbnailPublicID string            `json:"thumbnail_public_id,omitempty"`
	Stock             uint              `json:"stock"`
	Price             uint              `json:"price"`
	ExpirationDate    *time.Time        `json:"expiration_date,omitempty"`
	Category          *CategoryResponse `json:"category,omitempty"`
	CreatedAt         *time.Time        `json:"created_at,omitempty"`
	UpdatedAt         *time.Time        `json:"updated_at,omitempty"`
}
