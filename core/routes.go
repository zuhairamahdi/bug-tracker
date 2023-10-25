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
	var group = *api
	group.Get("/user", routes.GetAllUsers)
	group.Get("/api/user", routes.GetAllUsers)
	group.Post("/api/user", routes.CreateUser)
}

func setupBoardsRoute(api *fiber.Router) {
	var group = *api
	group.Get("/board", routes.GetAllBoards)
	group.Post("/board", routes.CreateBoard)
	group.Get("/board/:id", routes.GetBoard)
}
