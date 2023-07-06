package routes

import (
	"bugtracker/repos"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	users := repos.GetAllUsers()
	return c.Status(http.StatusOK).JSON(users)
}
