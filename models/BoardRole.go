package models


type BoardRole struct {
	Id int `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	BoardId int `json:"board_id"`
	Users []User `json:"users"`
	Name string `json:"name"`
	Description string `json:"description"`
}