package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadidrusalawi/gocare-api/internal/handler"
	"github.com/muhammadidrusalawi/gocare-api/internal/middleware"
	"github.com/muhammadidrusalawi/gocare-api/provider/mail"
)

func ApiRoute(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/auth/register", handler.RegisterHandler)
	api.Post("/auth/login", handler.LoginHandler)
	api.Post("/auth/logout", middleware.AuthMiddleware, handler.LogoutHandler)
	api.Get("/auth/profile", middleware.AuthMiddleware, handler.GetProfileHandler)

	api.Get("/admin/categories", middleware.AuthMiddleware, middleware.RoleMiddleware("admin"), handler.AdminGetAllCategoriesHandler)
	api.Post("/admin/categories", middleware.AuthMiddleware, middleware.RoleMiddleware("admin"), handler.AdminCreateCategoryHandler)
	api.Get("/admin/categories/:id", middleware.AuthMiddleware, middleware.RoleMiddleware("admin"), handler.AdminGetCategoryByIDHandler)
	api.Put("/admin/categories/:id", middleware.AuthMiddleware, middleware.RoleMiddleware("admin"), handler.AdminUpdateCategoryHandler)
	api.Delete("/admin/categories/:id", middleware.AuthMiddleware, middleware.RoleMiddleware("admin"), handler.AdminDeleteCategoryHandler)

	api.Get("/test/mail", func(c *fiber.Ctx) error {
		err := mail.Send(
			"cleaverrascal1@gmail.com",
			"TEST EMAIL",
			"<h1>Email berhasil dikirim 🚀</h1>",
		)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"success": true,
			"message": "email sent",
		})
	})

	api.Get("/admin/dashboard", middleware.AuthMiddleware, middleware.RoleMiddleware("admin"), handler.AdminDashboardHandler)
}
