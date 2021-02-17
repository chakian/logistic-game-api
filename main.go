package main

import (
	"github.com/chakian/logistic-game-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	setupRoutes(app)

	err := app.Listen(":8000")

	if err != nil {
		panic(err)
	}
}

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "This is base",
		})
	})

	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api endpoint",
		})
	})
	routes.TruckRoute(api.Group("/trucks"))
}
