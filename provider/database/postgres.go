package database

import (
	"fmt"
	"log"
	"os"

	"github.com/muhammadidrusalawi/gocare-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL_MODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	DB = db

	if err := DB.AutoMigrate(&model.User{}, &model.Category{}, &model.Product{}); err != nil {
		log.Fatal("failed to migrate:", err)
	}

	log.Println("Database connected")
}
