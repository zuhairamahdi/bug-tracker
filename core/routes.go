package core

import (
	"bugtracker/routes"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	setupUsersRoute(app)
	app.Listen(":3000")
}

func setupUsersRoute(app *fiber.App) {
	app.Get("/api/user/", routes.GetAllUsers)
}
