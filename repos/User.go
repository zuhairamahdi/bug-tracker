package repos

import (
	"bugtracker/ext"
	"bugtracker/models"
	"bugtracker/structs"
	"crypto/sha512"
	"encoding/hex"
	"errors"

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

func (r *userRepo) Login(username string, password string) (structs.User, error) {
	user := models.User{}
	responseUser := structs.User{}
	if err := r.storage.Find(&user).Where("username = ?", username).Error; err != nil {
		return responseUser, errors.New("Incorrect username or password")
	}
	if comparePass(password, user.Password, user.Salt) {
		responseUser = structs.User{
			Id:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}
		return responseUser, nil
	}
	return responseUser, errors.New("Incorrect username or password")
}

func (r *userRepo) Create(user structs.NewUser) error {
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
func (r *userRepo) Delete(id string) error {
	user := models.User{}
	if err := r.storage.Find(&user).Where("id =?", id).Error; err != nil {
		return err
	}
	if err := r.storage.Delete(&user).Error; err != nil {
		return err
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

func comparePass(password string, saltedPass string, salt string) bool {
	hashedPass := sha512.Sum512([]byte(password))
	requestedSaltedPass := hex.EncodeToString(hashedPass[:]) + salt
	saltedPassword := sha512.Sum512([]byte(requestedSaltedPass))

	if saltedPass != hex.EncodeToString(saltedPassword[:]) {
		return false
	}

	return true

}
