package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RoleMiddleware(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {

		claimsVal := c.Locals("claims")
		if claimsVal == nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}

		claims := claimsVal.(jwt.MapClaims)

		roleVal, ok := claims["role"]
		if !ok {
			return fiber.NewError(fiber.StatusForbidden, "Forbidden")
		}

		role := roleVal.(string)

		for _, allowed := range allowedRoles {
			if role == allowed {
				return c.Next()
			}
		}

		return fiber.NewError(fiber.StatusForbidden, "Forbidden")
	}
}
