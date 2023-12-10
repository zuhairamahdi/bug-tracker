package repos

import (
	"bugtracker/models"
	"bugtracker/structs"
	"errors"

	"gorm.io/gorm"
)

type boardUserRoleRepo struct {
	storage *gorm.DB
}

func newBoardUserRoleRepo(storage *gorm.DB) *boardUserRoleRepo {
	return &boardUserRoleRepo{storage: storage}
}

func (r *boardUserRoleRepo) AssignUserToBoardRole(user structs.User, role models.Role, board models.Board) error {
	//Check if user is already assigned to role
	boardUserRole := models.BoardUserRole{}
	if err := r.storage.Find(&boardUserRole).Where("user_id =? AND role_id =? AND board_id =?", user.Id, role.ID, board.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			boardUserRole = models.BoardUserRole{
				UserID:  user.Id,
				RoleID:  role.ID,
				BoardID: board.ID,
			}
			if err := r.storage.Create(&boardUserRole).Error; err != nil {
				return err
			}
		}
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
	//Get user roles from Role
	userModel := models.User{}
	if err := r.storage.Find(&user).Where("id =?", user.Id).Error; err != nil {
		return false, err
	}
	for _, role := range userModel.Roles {
		if role.Name == "admin" {
			return true, nil
		}
	}

	if len(boardUserRoles) > 0 {
		return true, nil
	}
	return false, nil
}
