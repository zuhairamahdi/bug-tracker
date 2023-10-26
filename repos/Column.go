package repos

import (
	"bugtracker/models"
	"bugtracker/structs"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type columnRepo struct {
	storage *gorm.DB
}

func newColumnRepo(storage *gorm.DB) *columnRepo {
	return &columnRepo{storage: storage}
}

func (r *columnRepo) GetAllByBoard(boardId string) ([]models.Column, error) {
	columns := []models.Column{}
	err := r.storage.Where("id = ?", boardId).Find(&columns).Error
	return columns, err
}

func (r *columnRepo) Create(column structs.Column) error {
	newColumn := models.Column{
		ID:         ulid.Make().String(),
		Title:      column.Title,
		Desciption: column.Desciption,
		Color:      column.Color,
		BoardID:    column.BoardID,
	}
	if query := r.storage.Create(&newColumn); query.Error != nil {
		return query.Error
	}
	return nil
}

func (r *columnRepo) GetByID(id string) (models.Column, error) {
	column := models.Column{}
	err := r.storage.Where("id =?", id).First(&column).Error
	return column, err
}

func (r *columnRepo) Update(column structs.Column) error {
	columnToUpdate := models.Column{
		ID:         column.ID,
		Title:      column.Title,
		Desciption: column.Desciption,
		Color:      column.Color,
		BoardID:    column.BoardID,
	}
	if query := r.storage.Save(&columnToUpdate); query.Error != nil {
		return query.Error
	}
	return nil
}
