package main

import (
	"go-pos-stores/app/routes"
	"go-pos-stores/app/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	utils.Init()

	routes.RoutesRegistry(app)

	app.Listen(":3000")

}
