package repos

import (
	"bugtracker/ext"
	"bugtracker/models"
	"bugtracker/structs"
	"crypto/sha512"
	"encoding/hex"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type userRepo struct {
	storage *gorm.DB
}

func newUserRepo(storage *gorm.DB) *userRepo {
	return &userRepo{
		storage: storage,
	}
}

// type userRepository interface {
// 	GetAll() []models.User
// }

func (r *userRepo) GetAll() []structs.User {
	Users := []models.User{}
	r.storage.Find(&Users)
	var allUsers []structs.User
	for _, user := range Users {
		allUsers = append(allUsers, structs.User{
			Id:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		})
	}
	return allUsers
}

func (r *userRepo) CreateNewUser(user structs.NewUser) error {
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
	if query := r.storage.Create(&newUser); query.Error != nil {
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
