package repos

import (
	"bugtracker/models"
	"bugtracker/storage"
	"bugtracker/structs"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type boardRepo struct {
	storage *gorm.DB
}

func newBoardRepo(storage *gorm.DB) *boardRepo {

	return &boardRepo{
		storage: storage,
	}
}

func (r *boardRepo) GetAll() ([]models.Board, error) {
	boards := []models.Board{}
	err := r.storage.Find(&boards).Error
	return boards, err
}

func (r *boardRepo) Create(board structs.Board) error {
	newBoard := models.Board{
		ID:         ulid.Make().String(),
		Title:      board.Title,
		Desciption: board.Description,
		UserID:     board.UserId,
	}
	if query := r.storage.Create(&newBoard); query.Error != nil {
		return query.Error
	}
	return nil
}

func (r *boardRepo) Get(uilid string) models.Board {
	board := models.Board{}
	r.storage.First(&board, "id = ?", uilid)
	return board
}

func (r *boardRepo) Delete(ulid string) error {
	board := r.Get(ulid)
	if query := storage.ApplicationDB.Delete(&board); query.Error != nil {
		return query.Error
	}
	return nil
}
