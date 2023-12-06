package repos

import (
	"bugtracker/models"

	"gorm.io/gorm"
)

type boardUserRoleRepo struct {
	storage *gorm.DB
}

func newBoardUserRoleRepo(storage *gorm.DB) *boardUserRoleRepo {
	return &boardUserRoleRepo{storage: storage}
}

func (r *boardUserRoleRepo) AssignUserToBoardRole(user models.User, role models.Role, board models.Board) error {
	boardUserRole := models.BoardUserRole{
		UserID:  user.ID,
		RoleID:  role.Id,
		BoardID: board.ID,
	}
	if query := r.storage.Create(&boardUserRole); query.Error != nil {
		return query.Error
	}
	return nil
}

func (r *boardUserRoleRepo) UnassignUserFromBoardRole(userRepo models.User, role models.Role, board models.Board) error {
	boardUserRole := models.BoardUserRole{
		UserID:  userRepo.ID,
		RoleID:  role.Id,
		BoardID: board.ID,
	}
	if query := r.storage.Delete(&boardUserRole); query.Error != nil {
		return query.Error
	}
	return nil
}
