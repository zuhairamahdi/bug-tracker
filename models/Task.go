package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	Status    string         `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	UserID    uint           `json:"user_id"`
	BoardID   uint           `json:"board_id"`
}
