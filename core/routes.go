package core

import (
	"bugtracker/middleware"
	"bugtracker/routes"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	apiGroup := app.Group("/api/v1")
	setupUsersRoute(&apiGroup)
	setupBoardsRoute(&apiGroup)
	setupPublicRoutes(&apiGroup)
	setupRolesRoute(&apiGroup)
	setupBoardUserRolesRoute(&apiGroup)
	app.Listen(":3000")
}

func setupPublicRoutes(api *fiber.Router) {
	var group = (*api).Group("/p")
	group.Post("/login", routes.Login)
}

func setupUsersRoute(api *fiber.Router) {
	var group = (*api).Group("/user").Use(middleware.JWTmiddleware)

	group.Get("/", routes.GetAllUsers)
	group.Post("/", routes.CreateUser)
	group.Put("/:id", routes.UpdateUser)
	group.Delete("/:id", routes.DeleteUser)
	group.Put("/:id/activate", routes.ActivateUser)
	group.Put("/:id/deactivate", routes.DeactivateUser)
}

func setupBoardsRoute(api *fiber.Router) {
	var group = (*api).Group("/board").Use(middleware.JWTmiddleware)
	group.Get("/", routes.GetAllBoards)
	group.Post("/", routes.CreateBoard)
	group.Get("/:id", routes.GetBoard)
}

func setupRolesRoute(api *fiber.Router) {
	var group = (*api).Group("/role").Use(middleware.JWTmiddleware)
	group.Get("/", routes.GetAllRoles)
	group.Post("/", routes.CreateRole)
	group.Get("/:id", routes.GetRole)
	group.Get("get-users-by-name/:name", routes.GetUsersByRoleName)
}

// Setup Board User Roles routes
func setupBoardUserRolesRoute(api *fiber.Router) {
	var group = (*api).Group("/board-user-role").Use(middleware.JWTmiddleware)
	group.Get("/:id", routes.GetAllBoardUserRoles)
	group.Get("/:board_id/:role_id", routes.HasAccessToBoardUsers)
	group.Post("/assign/:board_id/:role_id", routes.AssignBoardUserRoles)
	group.Post("/unassign/:board_id/:role_id", routes.UnassignUserFromBoardRole)
}
