package routes

import (
	"bugtracker/repos"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetAllBoards(c *fiber.Ctx) error {
	boards := repos.GetAllBoards()
	return c.Status(http.StatusOK).JSON(boards)
}
