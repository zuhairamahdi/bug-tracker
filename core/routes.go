package core

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	app.Listen(":3000")
}
