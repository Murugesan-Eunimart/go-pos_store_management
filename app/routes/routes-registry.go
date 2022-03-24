package routes

import "github.com/gofiber/fiber/v2"

func RoutesRegistry(app *fiber.App) {
	api := app.Group("/api/v1/pos_stores_management")

	PosCustomersRoutes(api)
}
