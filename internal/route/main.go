package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadidrusalawi/gocare-api/internal/handler"
	"github.com/muhammadidrusalawi/gocare-api/internal/middleware"
)

func ApiRoute(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/auth/register", handler.RegisterHandler)
	api.Post("/auth/verify-email", handler.VerifyEmailHandler)
	api.Post("/auth/login", handler.LoginHandler)
	api.Post("/auth/logout", middleware.AuthMiddleware, handler.LogoutHandler)
	api.Get("/auth/profile", middleware.AuthMiddleware, handler.GetProfileHandler)

	api.Get("/admin/categories", middleware.AuthMiddleware, middleware.RoleMiddleware("admin"), handler.AdminGetAllCategoriesHandler)
	api.Post("/admin/categories", middleware.AuthMiddleware, middleware.RoleMiddleware("admin"), handler.AdminCreateCategoryHandler)
	api.Get("/admin/categories/:id", middleware.AuthMiddleware, middleware.RoleMiddleware("admin"), handler.AdminGetCategoryByIDHandler)
	api.Put("/admin/categories/:id", middleware.AuthMiddleware, middleware.RoleMiddleware("admin"), handler.AdminUpdateCategoryHandler)
	api.Delete("/admin/categories/:id", middleware.AuthMiddleware, middleware.RoleMiddleware("admin"), handler.AdminDeleteCategoryHandler)

	api.Get("/admin/products", middleware.AuthMiddleware, middleware.RoleMiddleware("admin"), handler.AdminGetAllProductsHandler)
	api.Post("/admin/products", middleware.AuthMiddleware, middleware.RoleMiddleware("admin"), handler.AdminCreateProductHandler)
	api.Get("/admin/products/:id", middleware.AuthMiddleware, middleware.RoleMiddleware("admin"), handler.AdminGetProductByIDHandler)
	api.Put("/admin/products/:id", middleware.AuthMiddleware, middleware.RoleMiddleware("admin"), handler.AdminUpdateProductHandler)
	api.Delete("/admin/products/:id", middleware.AuthMiddleware, middleware.RoleMiddleware("admin"), handler.AdminDeleteProductHandler)

	api.Post("/upload", middleware.AuthMiddleware, middleware.RoleMiddleware("admin"), handler.ImageUploadHandler)
	api.Delete("/upload", middleware.AuthMiddleware, middleware.RoleMiddleware("admin"), handler.ImageDeleteHandler)

	api.Get("/admin/dashboard", middleware.AuthMiddleware, middleware.RoleMiddleware("admin"), handler.AdminDashboardHandler)
}
