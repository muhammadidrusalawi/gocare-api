package model

import "time"

type Address struct {
	ID           string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID       string `gorm:"not null;index"`
	ReceiverName string `gorm:"size:255;not null"`
	Phone        string `gorm:"size:20;not null"`
	Province     string `gorm:"size:255;not null"`
	City         string `gorm:"size:255;not null"`
	District     string `gorm:"size:255;not null"`
	PostalCode   string `gorm:"size:10;not null"`
	FullAddress  string `gorm:"type:text;not null"`
	IsDefault    bool   `gorm:"default:false;index"`
	Label        string `gorm:"size:50;not null;default:'home';check:label IN ('home','office','apartment','other')"`

	CreatedAt time.Time
	UpdatedAt time.Time

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
