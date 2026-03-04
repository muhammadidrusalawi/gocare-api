package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadidrusalawi/gocare-api/provider/auth"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	const prefix = "Bearer "
	if len(authHeader) <= len(prefix) || authHeader[:len(prefix)] != prefix {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token format")
	}

	tokenString := authHeader[len(prefix):]

	claims, err := auth.ParseToken(tokenString)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	jti := claims["jti"].(string)

	revoked, err := auth.IsTokenRevoked(jti)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Server error")
	}

	if revoked {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthenticated")
	}

	c.Locals("token", tokenString)
	c.Locals("claims", claims)

	return c.Next()
}
