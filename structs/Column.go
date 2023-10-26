package structs

import (
	"time"
)

type Column struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Desciption string `json:"description"`
	Color      string `json:"color"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
	BoardID    string
}
