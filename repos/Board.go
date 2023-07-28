package repos

import (
	"bugtracker/models"
	"bugtracker/storage"
	"bugtracker/structs"

	"github.com/oklog/ulid/v2"
)

func GetAllBoards() []models.Board {
	boards := []models.Board{}
	storage.ApplicationDB.Find(&boards)
	return boards
}

func CreateBoard(board structs.Board) error {
	newBoard := models.Board{
		ID:         ulid.Make().String(),
		Title:      board.Title,
		Desciption: board.Description,
		UserID:     board.UserId,
	}
	if query := storage.ApplicationDB.Create(&newBoard); query.Error != nil {
		return query.Error
	}
	return nil
}
func GetBoard(uilid string) models.Board {
	board := models.Board{}
	storage.ApplicationDB.First(&board, "id = ?", uilid)
	return board
}
