package models

type BoardUserRole struct {
	BoardID string `json:"board_id"`
	UserID  string `json:"user_id"`
	RoleID  uint   `json:"role_id"`
}
