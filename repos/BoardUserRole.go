package repos

import (
	"bugtracker/models"
	"bugtracker/structs"

	"gorm.io/gorm"
)

type boardUserRoleRepo struct {
	storage *gorm.DB
}

func newBoardUserRoleRepo(storage *gorm.DB) *boardUserRoleRepo {
	return &boardUserRoleRepo{storage: storage}
}

func (r *boardUserRoleRepo) AssignUserToBoardRole(user structs.User, role models.Role, board models.Board) error {
	boardUserRole := models.BoardUserRole{
		UserID:  user.Id,
		RoleID:  role.ID,
		BoardID: board.ID,
	}
	if query := r.storage.Create(&boardUserRole); query.Error != nil {
		return query.Error
	}
	return nil
}

func (r *boardUserRoleRepo) UnassignUserFromBoardRole(user structs.User, role models.Role, board models.Board) error {
	boardUserRole := models.BoardUserRole{
		UserID:  user.Id,
		RoleID:  role.ID,
		BoardID: board.ID,
	}
	if query := r.storage.Delete(&boardUserRole); query.Error != nil {
		return query.Error
	}
	return nil
}

func (r *boardUserRoleRepo) GetBoardUserRoles(board models.Board) ([]models.BoardUserRole, error) {
	var boardUserRoles []models.BoardUserRole
	if query := r.storage.Where("board_id =?", board.ID).Find(&boardUserRoles); query.Error != nil {
		return nil, query.Error
	}
	return boardUserRoles, nil
}

func (r *boardUserRoleRepo) IsUserHasRoleForBoard(user structs.User, role models.Role, board models.Board) (bool, error) {
	var boardUserRoles []models.BoardUserRole
	if query := r.storage.Where("board_id =? and user_id =? and role_id =?", board.ID, user.Id, role.ID).Find(&boardUserRoles); query.Error != nil {
		return false, query.Error
	}
	if len(boardUserRoles) > 0 {
		return true, nil
	}
	return false, nil
}
