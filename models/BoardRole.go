package models


type BoardRole struct {
	Id int `json:"id"`
	BoardId int `json:"board_id"`
	Users []User `json:"users"`
	Name string `json:"name"`
	Description string `json:"description"`
}