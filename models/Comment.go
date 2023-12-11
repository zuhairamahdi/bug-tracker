package models

type Comment struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	TaskID    uint   `json:"task_id"`
	UserID    uint   `json:"user_id"`
}
