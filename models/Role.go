package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	Id        string `json:"id" gorm:"primaryKey;autoIncrement:false"`
	Name      string `json:"name" gorm:"index;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Users     []User         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
