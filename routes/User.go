package routes

import (
	"bugtracker/repos"
	"bugtracker/structs"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	users := repos.GetAllUsers()
	return c.Status(http.StatusOK).JSON(users)
}
func CreateUser(c *fiber.Ctx) error {
	user := structs.NewUser{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	if err := repos.CreateNewUser(user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(structs.ErrorResponse{ErrorCode: "ERR01", Message: err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(user)
}
