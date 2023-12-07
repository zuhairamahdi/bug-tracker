package routes

import (
	appcontext "bugtracker/app-context"
	"bugtracker/repos"
	"bugtracker/structs"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllBoardUserRoles(c *fiber.Ctx) error {
	ulid := c.Params("id")
	board := repos.Repos.BoardRepository.Get(ulid)
	boardUserRoles, err := repos.Repos.BoardUserRoleRepository.GetBoardUserRoles(board)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(boardUserRoles)
}

func HasAccessToBoardUsers(c *fiber.Ctx) error {
	user, err := appcontext.GetUserData(c)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(structs.ErrorResponse{ErrorCode: "ERR0", Message: err.Error()})
	}
	boardId := c.Params("board_id")
	roleId, _ := strconv.ParseUint(c.Params("role_id"), 10, 64)

	board := repos.Repos.BoardRepository.Get(boardId)
	role, _ := repos.Repos.RoleRepository.FindById(uint(roleId))
	hasAccess, err := repos.Repos.BoardUserRoleRepository.IsUserHasRoleForBoard(user, role, board)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(hasAccess)
}

func AssignBoardUserRoles(c *fiber.Ctx) error {
	user, err := appcontext.GetUserData(c)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(structs.ErrorResponse{ErrorCode: "ERR0", Message: err.Error()})
	}
	boardId := c.Params("board_id")
	roleId, _ := strconv.ParseUint(c.Params("role_id"), 10, 64)

	board := repos.Repos.BoardRepository.Get(boardId)
	role, _ := repos.Repos.RoleRepository.FindById(uint(roleId))
	err = repos.Repos.BoardUserRoleRepository.AssignUserToBoardRole(user, role, board)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(struct{}{})
}

func UnassignUserFromBoardRole(c *fiber.Ctx) error {
	user, err := appcontext.GetUserData(c)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(structs.ErrorResponse{ErrorCode: "ERR0", Message: err.Error()})
	}
	boardId := c.Params("board_id")
	roleId, _ := strconv.ParseUint(c.Params("role_id"), 10, 64)

	board := repos.Repos.BoardRepository.Get(boardId)
	role, _ := repos.Repos.RoleRepository.FindById(uint(roleId))
	err = repos.Repos.BoardUserRoleRepository.UnassignUserFromBoardRole(user, role, board)
	if err != nil {
		return err
	}
	return c.SendStatus(http.StatusOK)
}
