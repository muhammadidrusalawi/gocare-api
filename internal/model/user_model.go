package model

import (
	"time"
)

type User struct {
	ID       string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name     string `gorm:"size:255;not null"`
	Email    string `gorm:"size:100;not null;unique"`
	Password string `gorm:"not null"`
	Role     string `gorm:"size:50;not null;default:'customer';check:role IN ('admin','customer')"`

	VerifiedAt *time.Time `gorm:"default:null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
