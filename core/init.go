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
	SetupRoutes(app)
}
