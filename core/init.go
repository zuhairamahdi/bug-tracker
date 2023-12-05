package core

import (
	"bugtracker/models"
	"bugtracker/repos"
	"bugtracker/storage"

	"github.com/gofiber/fiber/v2"
)

func Init() {
	storage.InitializeDB()
	models.Migrate()
	app := fiber.New()
	repos.InitializeRepositories()
	repos.Repos.RoleRepository.MigrateInitialRoles()
	repos.Repos.UserRepository.MigrateInitialUser()
	SetupRoutes(app)
}
