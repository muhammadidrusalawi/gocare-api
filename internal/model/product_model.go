package model

import (
	"time"
)

type Product struct {
	ID                     string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name                   string `gorm:"size:255;not null"`
	Slug                   string `gorm:"size:255;not null;unique"`
	CategoryID             string `gorm:"not null;index"`
	Description            string `gorm:"type:text"`
	Thumbnail              string `gorm:"type:text"`
	ThumbnailPublicID      string `gorm:"type:text"`
	Stock                  uint
	Price                  uint
	IsPrescriptionRequired bool       `gorm:"default:false"`
	ExpirationDate         *time.Time `gorm:"type:date;default:null"`

	CreatedAt time.Time
	UpdatedAt time.Time

	Category Category `gorm:"foreignKey:CategoryID"`
}
