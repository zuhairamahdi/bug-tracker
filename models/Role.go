package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name      string `json:"name" gorm:"index;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
