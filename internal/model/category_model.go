package model

import (
	"time"
)

type Category struct {
	ID          string  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string  `gorm:"size:255;not null;unique"`
	Slug        string  `gorm:"size:255;not null;unique"`
	Description *string `gorm:"type:text"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
