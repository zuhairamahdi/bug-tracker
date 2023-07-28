package models

import (
	"time"

	"gorm.io/gorm"
)

type Board struct {
	ID         string `json:"id" gorm:"primaryKey;autoIncrement:false"`
	Title      string `json:"title" gorm:"index"`
	Desciption string `json:"description"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	UserID     string
	Coulmns    []Column `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
