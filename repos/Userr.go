package repos

import (
	"bugtracker/models"
	"bugtracker/storage"
	"fmt"
)

func GetAllUsers() []models.User {
	Users := []models.User{}
	storage.ApplicationDB.Find(&Users)
	fmt.Println(Users)
	return Users
}
