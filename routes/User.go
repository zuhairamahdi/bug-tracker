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
	if err := repos.Repos.UserRepository.Create(user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(structs.ErrorResponse{ErrorCode: "ERR01", Message: err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := structs.NewUser{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	if err := repos.Repos.UserRepository.Update(id, user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(structs.ErrorResponse{ErrorCode: "ERR01", Message: err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(user)
}

// Add user to role
func AddUserToRole(c *fiber.Ctx) error {
	id := c.Params("id")
	role := c.Params("role")
	if err := repos.Repos.UserRepository.AddUserToRole(id, role); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(structs.ErrorResponse{ErrorCode: "ERR01", Message: err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(role)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := repos.Repos.UserRepository.Delete(id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(structs.ErrorResponse{ErrorCode: "ERR01", Message: err.Error()})
	}
	return c.Status(http.StatusNoContent).JSON(nil)
}
