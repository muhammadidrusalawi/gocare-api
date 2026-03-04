package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadidrusalawi/gocare-api/internal/model"
	"github.com/muhammadidrusalawi/gocare-api/provider/auth"
	"github.com/muhammadidrusalawi/gocare-api/provider/database"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(name, email, password string) (*model.User, error) {
	var existing model.User
	if err := database.DB.Where("email = ?", email).First(&existing).Error; err == nil {
		return nil, fiber.NewError(fiber.StatusConflict, "Email already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Name:     name,
		Email:    email,
		Password: string(hashed),
		Role:     "customer",
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func LoginUser(email, password string) (*model.User, string, error) {
	var user model.User

	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, "", fiber.NewError(fiber.StatusBadRequest, "Email or password is incorrect")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", fiber.NewError(fiber.StatusBadRequest, "Email or password is incorrect")
	}

	if user.VerifiedAt == nil {
		return nil, "", fiber.NewError(fiber.StatusForbidden, "User not found")
	}

	token, err := auth.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, "", err
	}

	return &user, token, nil
}

func ProfileUser(userID string) (*model.User, error) {
	var user model.User

	if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func LogoutUser(token string) error {
	return auth.RevokeToken(token)
}
