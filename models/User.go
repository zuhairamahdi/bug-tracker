package models

import (
	"database/sql"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type User struct {
	ID          ulid.ULID `json:"id" gorm:"primaryKey;autoIncrement:false"`
	Username    string    `json:"username" gorm:"index"`
	Email       string    `json:"email" gorm:"index"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Password    string
	Salt        string
	ActivatedAt sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
