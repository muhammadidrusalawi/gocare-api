package middleware

import (
	"log"
	"os"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/muhammadidrusalawi/gocare-api/internal/helper"
)

func ErrorMiddleware() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		appEnv := strings.ToLower(os.Getenv("APP_ENV"))

		appDebug, parseErr := strconv.ParseBool(os.Getenv("APP_DEBUG"))
		if parseErr != nil {
			appDebug = false
		}

		if e, ok := err.(*fiber.Error); ok {
			return c.Status(e.Code).JSON(
				helper.ApiError(e.Message),
			)
		}

		log.Printf("ERROR: %v\n%s", err, debug.Stack())

		if appEnv == "development" && appDebug {
			return c.Status(fiber.StatusInternalServerError).JSON(
				helper.ApiError(err.Error() + "\n\n" + string(debug.Stack())),
			)
		}

		return c.Status(fiber.StatusInternalServerError).JSON(
			helper.ApiError("internal server error"),
		)
	}
}
