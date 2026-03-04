package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/muhammadidrusalawi/gocare-api/internal/middleware"
	"github.com/muhammadidrusalawi/gocare-api/internal/route"
	"github.com/muhammadidrusalawi/gocare-api/provider/cache"
	"github.com/muhammadidrusalawi/gocare-api/provider/database"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	database.InitDB()
	cache.ConnectRedis()

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorMiddleware(),
	})

	app.Use(logger.New(logger.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Method() == fiber.MethodOptions
		},
	}))

	app.Use(cors.New())

	route.ApiRoute(app)

	log.Fatal(app.Listen(":3000"))
}
