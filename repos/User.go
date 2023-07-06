package repos

import (
	"bugtracker/ext"
	"bugtracker/models"
	"bugtracker/storage"
	"bugtracker/structs"
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"github.com/oklog/ulid/v2"
)

const salter = "BF$#DAVcvdksj@@31"

func GetAllUsers() []models.User {
	Users := []models.User{}
	storage.ApplicationDB.Find(&Users)
	fmt.Println(Users)
	return Users
}

func CreateNewUser(user structs.NewUser) error {
	salt, pass := createSaltedPass(user.Password)
	newUser := models.User{
		ID:        ulid.Make().String(),
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  pass,
		Salt:      salt,
	}
	if query := storage.ApplicationDB.Create(&newUser); query.Error != nil {
		return query.Error
	}
	return nil
}

func createSaltedPass(password string) (string, string) {
	hashedPass := sha512.Sum512([]byte(password))
	hashedSalt := sha512.Sum512([]byte(ext.RandStringBytes(10)))
	saltedPass := hex.EncodeToString(hashedPass[:]) + hex.EncodeToString(hashedSalt[:])
	SaltedPassword := sha512.Sum512([]byte(saltedPass))
	return hex.EncodeToString(hashedSalt[:]), hex.EncodeToString(SaltedPassword[:])

}
