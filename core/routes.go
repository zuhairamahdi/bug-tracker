package core

import (
	"bugtracker/routes"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	setupUsersRoute(app)
	setupBoardsRoute(app)
	app.Listen(":3000")
}

func setupUsersRoute(app *fiber.App) {
	app.Get("/api/user/", routes.GetAllUsers)
	app.Post("/api/user", routes.CreateUser)
}

func setupBoardsRoute(app *fiber.App) {

}
