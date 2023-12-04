package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          string `json:"id" gorm:"primaryKey;autoIncrement:false"`
	Username    string `json:"username" gorm:"index;unique"`
	Email       string `json:"email" gorm:"index;unique"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Password    string
	Salt        string
	Active      bool `json:"active"`
	ActivatedAt sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Boards      []Board        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Roles       []Role         `gorm:"many2many:user_roles,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
