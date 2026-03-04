package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/muhammadidrusalawi/gocare-api/internal/model"
	"github.com/muhammadidrusalawi/gocare-api/provider/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()
	database.InitDB()

	db := database.DB

	tx := db.Begin()

	if err := seedCategories(tx); err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	if err := seedUser(tx); err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()

	log.Println("Seeding completed 🚀")
}

func seedCategories(db *gorm.DB) error {
	categories := []model.Category{
		{Name: "Obat Bebas", Slug: "obat-bebas"},
		{Name: "Vitamin & Suplemen", Slug: "vitamin-suplemen"},
		{Name: "Peralatan Medis", Slug: "peralatan-medis"},
		{Name: "Produk Bayi", Slug: "produk-bayi"},
		{Name: "P3K", Slug: "p3k"},
		{Name: "Obat Herbal", Slug: "obat-herbal"},
		{Name: "Alat Bantu Dengar", Slug: "alat-bantu-dengar"},
		{Name: "Kebutuhan Harian", Slug: "kebutuhan-harian"},
	}

	for _, c := range categories {
		var existing model.Category

		err := db.Where("slug = ?", c.Slug).First(&existing).Error
		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&c).Error; err != nil {
				return err
			}
			log.Println("Seeded:", c.Name)
		}
	}

	return nil
}

func seedUser(db *gorm.DB) error {
	now := time.Now()

	users := []model.User{
		{
			Name:       "Admin Gocare",
			Email:      "admin@example.com",
			Password:   "12345678",
			Role:       "admin",
			VerifiedAt: &now,
		},
		{
			Name:       "Customer Gocare",
			Email:      "customer@example.com",
			Password:   "12345678",
			Role:       "customer",
			VerifiedAt: &now,
		},
	}

	for _, u := range users {
		var existing model.User

		err := db.Where("email = ?", u.Email).First(&existing).Error
		if err == gorm.ErrRecordNotFound {

			hashedPassword, err := bcrypt.GenerateFromPassword(
				[]byte(u.Password),
				bcrypt.DefaultCost,
			)
			if err != nil {
				return err
			}

			u.Password = string(hashedPassword)

			if err := db.Create(&u).Error; err != nil {
				return err
			}

			log.Println("Seeded user:", u.Email)
		}
	}

	return nil
}
