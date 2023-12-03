package routes

import (
	"bugtracker/repos"
	"bugtracker/structs"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	users := repos.Repos.UserRepository.GetAll()
	return c.Status(http.StatusOK).JSON(users)
}
func CreateUser(c *fiber.Ctx) error {
	user := structs.NewUser{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	if err := repos.Repos.UserRepository.CreateNewUser(user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(structs.ErrorResponse{ErrorCode: "ERR01", Message: err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := repos.Repos.UserRepository.DeleteUser(id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(structs.ErrorResponse{ErrorCode: "ERR01", Message: err.Error()})
	}
	return c.Status(http.StatusNoContent).JSON(nil)
}
