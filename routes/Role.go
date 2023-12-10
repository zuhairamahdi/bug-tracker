package routes

import (
	"bugtracker/models"
	"bugtracker/repos"
	"bugtracker/structs"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllRoles(c *fiber.Ctx) error {
	roles, err := repos.Repos.RoleRepository.GetAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(structs.ErrorResponse{ErrorCode: "ERR01", Message: err.Error()})
	}
	return c.Status(http.StatusOK).JSON(roles)
}

func CreateRole(c *fiber.Ctx) error {
	role := models.Role{}
	if err := c.BodyParser(&role); err != nil {
		return err
	}
	if err := repos.Repos.RoleRepository.Create(role); err != nil {
		return err
	}
	return c.Status(http.StatusCreated).JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	role := models.Role{}
	if err := c.BodyParser(&role); err != nil {
		return err
	}
	if err := repos.Repos.RoleRepository.Update(role); err != nil {
		return err
	}
	return c.Status(http.StatusCreated).JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {
	roleId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(structs.ErrorResponse{ErrorCode: "ERR0", Message: err.Error()})
	}
	if err := repos.Repos.RoleRepository.DeleteById(uint(roleId)); err != nil {
		return err
	}
	return c.Status(http.StatusCreated).JSON(roleId)
}

func GetRole(c *fiber.Ctx) error {
	roleId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(structs.ErrorResponse{ErrorCode: "ERR0", Message: err.Error()})
	}
	role, err := repos.Repos.RoleRepository.FindById(uint(roleId))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(structs.ErrorResponse{ErrorCode: "ERR01", Message: err.Error()})
	}
	return c.Status(http.StatusOK).JSON(role)
}

func GetUsersByRoleName(c *fiber.Ctx) error {
	roleName := c.Params("name")
	users, err := repos.Repos.RoleRepository.FindUsersByRoleName(roleName)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(structs.ErrorResponse{ErrorCode: "ERR01", Message: err.Error()})
	}
	return c.Status(http.StatusOK).JSON(users)
}
