package core

import (
	"bugtracker/routes"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	apiGroup := app.Group("/api/v1")
	setupUsersRoute(&apiGroup)
	setupBoardsRoute(&apiGroup)
	app.Listen(":3000")
}

func setupUsersRoute(api *fiber.Router) {
	var group = (*api).Group("/user")
	group.Get("/", routes.GetAllUsers)
	group.Post("/", routes.CreateUser)
}

func setupBoardsRoute(api *fiber.Router) {
	var group = (*api).Group("/board")
	group.Get("/", routes.GetAllBoards)
	group.Post("/", routes.CreateBoard)
	group.Get("/:id", routes.GetBoard)
}
