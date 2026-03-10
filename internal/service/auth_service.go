package service

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/muhammadidrusalawi/gocare-api/internal/model"
	"github.com/muhammadidrusalawi/gocare-api/internal/repository"
	"github.com/muhammadidrusalawi/gocare-api/internal/request"
	"github.com/muhammadidrusalawi/gocare-api/provider/auth"
	"github.com/muhammadidrusalawi/gocare-api/provider/cache"
	"github.com/muhammadidrusalawi/gocare-api/provider/database"
	"github.com/muhammadidrusalawi/gocare-api/provider/mail"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(req request.RegisterRequest) error {
	userRepo := repository.NewUserRepository(database.DB)
	userExist, err := userRepo.FindByEmail(req.Email)
	if err == nil && userExist != nil {
		return fiber.NewError(fiber.StatusConflict, "Email already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashed),
		Role:     "customer",
	}

	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	token := uuid.NewString()
	rateKey := "register:email:" + user.Email

	ok, _ := cache.Client.SetNX(
		cache.Ctx,
		rateKey,
		1,
		15*time.Minute,
	).Result()

	if !ok {
		return fiber.NewError(fiber.StatusTooManyRequests, "Verification link already sent. Please wait 15 minutes")
	}

	verifyKey := "register:token:" + token

	err = cache.Client.Set(
		cache.Ctx,
		verifyKey,
		data,
		15*time.Minute,
	).Err()

	if err != nil {
		return err
	}

	clientURL := os.Getenv("CLIENT_URL")

	verifyLink := fmt.Sprintf(
		"%s/auth/verify-email?token=%s",
		clientURL,
		token,
	)

	emailBody := fmt.Sprintf(`
		<h1>Verify your email</h1>
		<p>Click the link below to verify your account:</p>
		<a href="%s">%s</a>
	`, verifyLink, verifyLink)

	go func(email, subject, body string) {
		_ = mail.Send(email, subject, body)
	}(user.Email, "Email Verification", emailBody)

	return nil
}

func VerifyEmail(token string) (*model.User, string, error) {
	userRepo := repository.NewUserRepository(database.DB)
	key := "register:token:" + token

	data, err := cache.Client.Get(cache.Ctx, key).Result()
	if err != nil {
		return nil, "", fiber.NewError(fiber.StatusBadRequest, "Invalid or expired token")
	}

	defer cache.Client.Del(cache.Ctx, key)
	var user model.User

	if err := json.Unmarshal([]byte(data), &user); err != nil {
		return nil, "", err
	}

	existing, _ := userRepo.FindByEmail(user.Email)
	if existing != nil {
		return nil, "", fiber.NewError(fiber.StatusConflict, "Email already verified")
	}

	now := time.Now()
	user.VerifiedAt = &now

	if err := userRepo.Create(&user); err != nil {
		return nil, "", err
	}

	accessToken, err := auth.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, "", err
	}

	user.Password = ""

	return &user, accessToken, nil
}

func LoginUser(req request.LoginRequest) (*model.User, string, error) {
	userRepo := repository.NewUserRepository(database.DB)
	user, err := userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, "", fiber.NewError(fiber.StatusBadRequest, "Email or password is incorrect")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, "", fiber.NewError(fiber.StatusBadRequest, "Email or password is incorrect")
	}

	if user.VerifiedAt == nil {
		return nil, "", fiber.NewError(fiber.StatusForbidden, "User not found")
	}

	token, err := auth.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func GoogleAuthRedirect() (string, error) {
	state := uuid.NewString()
	url := auth.GoogleConfig().AuthCodeURL(state)

	return url, nil
}

func GoogleCallback(code string, state string) string {
	clientURL := os.Getenv("CLIENT_URL")

	redirectURL := fmt.Sprintf(
		"%s/auth/google-callback?code=%s&state=%s",
		clientURL,
		code,
		state,
	)

	return redirectURL
}

func GoogleAuthExchange(req request.GoogleExchange) (*model.User, string, error) {
	userRepo := repository.NewUserRepository(database.DB)

	oauthCode, err := auth.GoogleConfig().Exchange(context.Background(), req.Code)
	if err != nil {
		return nil, "", err
	}

	client := auth.GoogleConfig().Client(context.Background(), oauthCode)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	var googleUser struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&googleUser); err != nil {
		return nil, "", err
	}

	now := time.Now()
	randomPassword := uuid.NewString()

	hashed, err := bcrypt.GenerateFromPassword([]byte(randomPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", err
	}

	user := &model.User{
		Name:       googleUser.Name,
		Email:      googleUser.Email,
		Password:   string(hashed),
		Role:       "customer",
		VerifiedAt: &now,
	}

	if err := userRepo.Upsert(user); err != nil {
		return nil, "", err
	}

	token, err := auth.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func ProfileUser(userID string) (*model.User, error) {
	userRepo := repository.NewUserRepository(database.DB)
	user, err := userRepo.FindByID(userID)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	return user, nil
}

func LogoutUser(token string) error {
	return auth.RevokeToken(token)
}
