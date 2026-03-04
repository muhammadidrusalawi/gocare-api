package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/muhammadidrusalawi/gocare-api/internal/helper"
	"github.com/muhammadidrusalawi/gocare-api/internal/service"
)

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func RegisterHandler(c *fiber.Ctx) error {
	var req RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError("Invalid JSON"))
	}

	if err := helper.ValidateStruct(req); err != "" {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError(err))
	}

	user, err := service.RegisterUser(req.Name, req.Email, req.Password)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(helper.ApiSuccess("User registered successfully", fiber.Map{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"role":  user.Role,
	}))
}

func LoginHandler(c *fiber.Ctx) error {
	var req LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError("Invalid JSON"))
	}

	if err := helper.ValidateStruct(req); err != "" {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError(err))
	}

	user, token, err := service.LoginUser(req.Email, req.Password)
	if err != nil {
		return err
	}

	return c.JSON(helper.ApiSuccess("User logged in successfully", fiber.Map{
		"user": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
		"token": token,
	}))
}

func GetProfileHandler(c *fiber.Ctx) error {
	claims := c.Locals("claims").(jwt.MapClaims)
	userID := claims["user_id"].(string)

	user, err := service.ProfileUser(userID)
	if err != nil {
		return err
	}

	return c.JSON(helper.ApiSuccess("User profile successfully", fiber.Map{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"role":  user.Role,
	}))
}

func LogoutHandler(c *fiber.Ctx) error {
	tokenVal := c.Locals("token")
	if tokenVal == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	token := tokenVal.(string)

	if err := service.LogoutUser(token); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(helper.ApiSuccess("User logout successfully", nil))
}
