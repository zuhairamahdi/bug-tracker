package routes

import (
	appcontext "bugtracker/app-context"
	"bugtracker/repos"
	"bugtracker/structs"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetAllBoards(c *fiber.Ctx) error {
	boards, _ := repos.Repos.BoardRepository.GetAll()
	return c.Status(http.StatusOK).JSON(boards)
}

func CreateBoard(c *fiber.Ctx) error {
	board := structs.Board{}
	if err := c.BodyParser(&board); err != nil {
		return err
	}
	//get user details from appcontext
	user, err := appcontext.GetUserData(c)
	if err != nil {
		return err
	}
	board.UserId = user.Id
	//can user create board?
	canCreate, err := repos.Repos.RoleRepository.CanUserCreateBoard(user, board)
	if err != nil {
		return err
	}
	if !canCreate {
		return c.Status(http.StatusForbidden).JSON(map[string]string{"message": "You are not authorized to create this board"})
	}
	if err := repos.Repos.BoardRepository.Create(board); err != nil {
		return err
	}
	return c.Status(http.StatusCreated).JSON(board)

}

func GetBoard(c *fiber.Ctx) error {
	ulid := c.Params("id")
	board := repos.Repos.BoardRepository.Get(ulid)
	return c.Status(http.StatusOK).JSON(board)
}
func DeleteBoard(c *fiber.Ctx) error {
	ulid := c.Params("id")
	if err := repos.Repos.BoardRepository.Delete(ulid); err != nil {
		return err
	}
	return c.Status(http.StatusNoContent).JSON(nil)
}
