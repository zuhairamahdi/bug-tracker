package repos

import (
	"bugtracker/ext"
	"bugtracker/models"
	"bugtracker/structs"
	"crypto/rand"
	"errors"
	"log"
	"time"

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

func (r *columnRepo) CreateDefaultColumns(boardId string) error {
	//create 4 columns (pending, in progress, stuck and done)
	columns := []models.Column{
		{
			ID:         ulid.MustNew(ulid.Now(), nil).String(),
			Title:      "Pending",
			Desciption: "This column is used to show tasks that are not yet started",
			BoardID:    boardId,
			Color:      "#87a2c7",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         ulid.MustNew(ulid.Now(), nil).String(),
			Title:      "In Progress",
			Desciption: "This column is used to track the progress of the task",
			BoardID:    boardId,
			Color:      "#ffcc00",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         ulid.MustNew(ulid.Now(), nil).String(),
			Title:      "Stuck",
			Desciption: "This column is used to mark tasks that are stuck in the backlog",
			BoardID:    boardId,
			Color:      "#e03b24",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         ulid.MustNew(ulid.Now(), nil).String(),
			Title:      "Done",
			Desciption: "This column is used to mark tasks that have been completed",
			BoardID:    boardId,
			Color:      "#64a338",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	for i := 0; i < len(columns); i++ {
		entropy := make([]byte, 16)
		_, err := rand.Read(entropy)
		if err != nil {
			log.Fatal(err)
		}

		entropySource := rand.Reader
		if err != nil {
			log.Fatal(err)
		}
		id := ulid.MustNew(ulid.Now(), entropySource).String()
		columns[i].ID = id
	}

	err := r.storage.Create(&columns).Error
	return err
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

func (r *columnRepo) Delete(id string) error {
	if query := r.storage.Delete(&models.Column{ID: id}); query.Error != nil {
		return query.Error
	}
	return nil
}

func (r *columnRepo) UpdateTitle(id string, title string) error {
	if query := r.storage.Model(&models.Column{}).Where("id =?", id).Update("title", title); query.Error != nil {
		return query.Error
	}
	return nil
}

func (r *columnRepo) UpdateDesciption(id string, description string) error {
	if query := r.storage.Model(&models.Column{}).Where("id =?", id).Update("desciption", description); query.Error != nil {
		return query.Error
	}
	return nil
}

func (r *columnRepo) UpdateColor(id string, color string) error {
	// check if color is a valid hex color
	isValidColor := ext.IsColorHex(color)
	if !isValidColor {
		return errors.New("Invalid color")
	}
	if query := r.storage.Model(&models.Column{}).Where("id =?", id).Update("color", color); query.Error != nil {
		return query.Error
	}
	return nil
}
