package models

import (
	"time"

	"gorm.io/gorm"
)

type Column struct {
	ID         string `json:"id" gorm:"primaryKey;autoIncrement:false"`
	Title      string `json:"title" gorm:"index"`
	Desciption string `json:"description"`
	Color      string `json:"color"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	BoardID    string
}
