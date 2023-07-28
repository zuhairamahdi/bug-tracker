package routes

import (
	"bugtracker/repos"
	"bugtracker/structs"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetAllBoards(c *fiber.Ctx) error {
	boards := repos.GetAllBoards()
	return c.Status(http.StatusOK).JSON(boards)
}

func CreateBoard(c *fiber.Ctx) error {
	board := structs.Board{}
	if err := c.BodyParser(&board); err != nil {
		return err
	}
	if err := repos.CreateBoard(board); err != nil {
		return err
	}
	return c.Status(http.StatusCreated).JSON(board)
}

func GetBoard(c *fiber.Ctx) error {
	ulid := c.Params("id")

	board := repos.GetBoard(ulid)
	return c.Status(http.StatusOK).JSON(board)
}
