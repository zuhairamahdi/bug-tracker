package models

import "time"

type Invitation struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ExpiresAt time.Time `json:"expires_at"`
	Token     string    `json:"token"`
	BoardID   string    `json:"board_id"`
	UserID    string    `json:"user_id"`
	RoleID    string    `json:"role_id"`
	Status    string    `json:"status"`
}
