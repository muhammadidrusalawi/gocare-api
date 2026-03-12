package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/muhammadidrusalawi/gocare-api/internal/helper"
	"github.com/muhammadidrusalawi/gocare-api/internal/request"
	"github.com/muhammadidrusalawi/gocare-api/internal/response"
	"github.com/muhammadidrusalawi/gocare-api/internal/service"
)

func RegisterHandler(c *fiber.Ctx) error {
	var req request.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError("Invalid JSON"))
	}

	if err := helper.ValidateStruct(req); err != "" {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError(err))
	}

	err := service.RegisterUser(req)

	if err != nil {
		return err
	}

	msg := "The verification link has been sent to the " + req.Email

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess(msg, nil))
}

func VerifyEmailHandler(c *fiber.Ctx) error {
	var req request.VerifyEmailRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError("Invalid JSON"))
	}

	if err := helper.ValidateStruct(req); err != "" {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError(err))
	}

	user, accessToken, err := service.VerifyEmail(req.VerificationToken)
	if err != nil {
		return err
	}

	res := response.LoginResponse{
		User: response.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
		Token: accessToken,
	}

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("User logged in successfully", res))

}

func LoginHandler(c *fiber.Ctx) error {
	var req request.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError("Invalid JSON"))
	}

	if err := helper.ValidateStruct(req); err != "" {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError(err))
	}

	user, token, err := service.LoginUser(req)
	if err != nil {
		return err
	}

	res := response.LoginResponse{
		User: response.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
		Token: token,
	}

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("User logged in successfully", res))
}

func GoogleAuthRedirectHandler(c *fiber.Ctx) error {
	url, err := service.GoogleAuthRedirect()
	if err != nil {
		return err
	}

	return c.Redirect(url)
}

func GoogleAuthCallbackHandler(c *fiber.Ctx) error {
	code := c.Query("code")
	state := c.Query("state")

	redirectURL := service.GoogleCallback(code, state)

	return c.Redirect(redirectURL)
}

func GoogleAuthExchangeHandler(c *fiber.Ctx) error {
	var req request.GoogleExchange

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError("Invalid JSON"))
	}

	if err := helper.ValidateStruct(req); err != "" {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError(err))
	}

	user, token, err := service.GoogleAuthExchange(req)
	if err != nil {
		return err
	}

	res := response.LoginResponse{
		User: response.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
		Token: token,
	}

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("User logged in successfully", res))
}

func GetProfileHandler(c *fiber.Ctx) error {
	claims := c.Locals("claims").(jwt.MapClaims)
	userID := claims["user_id"].(string)

	user, err := service.ProfileUser(userID)
	if err != nil {
		return err
	}

	res := response.UserResponse{
		ID:         userID,
		Name:       user.Name,
		Email:      user.Email,
		Role:       user.Role,
		VerifiedAt: user.VerifiedAt,
		CreatedAt:  &user.CreatedAt,
		UpdatedAt:  &user.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("User profile retrieved successfully", res))
}

func ForgotPasswordHandler(c *fiber.Ctx) error {
	var req request.ForgotPasswordRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError("Invalid JSON"))
	}

	if err := helper.ValidateStruct(req); err != "" {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError(err))
	}

	err := service.ForgotPassword(req)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("A password reset link has been sent. Please check your inbox.", nil))
}

func ResetPasswordHandler(c *fiber.Ctx) error {
	var req request.ResetPasswordRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError("Invalid JSON"))
	}

	if err := helper.ValidateStruct(req); err != "" {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiError(err))
	}

	err := service.ResetPassword(req)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("Your password has been reset successfully.", nil))
}

func UpdateUserProfileHandler(c *fiber.Ctx) error {
	claims := c.Locals("claims").(jwt.MapClaims)
	userID := claims["user_id"].(string)

	var req request.UpdateUserProfileRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(helper.ApiError("Invalid JSON"))
	}

	if err := helper.ValidateStruct(req); err != "" {
		return c.Status(fiber.StatusBadRequest).
			JSON(helper.ApiError(err))
	}

	user, err := service.UpdateUserProfile(userID, req)
	if err != nil {
		return err
	}

	res := response.UserResponse{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Role:       user.Role,
		VerifiedAt: user.VerifiedAt,
		CreatedAt:  &user.CreatedAt,
		UpdatedAt:  &user.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("User profile updated successfully", res))
}

func LogoutHandler(c *fiber.Ctx) error {
	tokenVal := c.Locals("token")
	token := tokenVal.(string)

	if err := service.LogoutUser(token); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(helper.ApiSuccess("User logout successfully", nil))
}
