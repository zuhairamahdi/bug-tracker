package repos

import (
	"bugtracker/ext"
	"bugtracker/models"
	"bugtracker/structs"
	"crypto/sha512"
	"database/sql"
	"encoding/hex"
	"errors"
	"time"

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
func (r *userRepo) Update(id string, user structs.NewUser) error {
	userToUpdate := models.User{}
	if err := r.storage.Find(&userToUpdate).Where("id =?", id).Error; err != nil {
		return err
	}
	userToUpdate.Username = user.Username
	userToUpdate.Email = user.Email
	userToUpdate.FirstName = user.FirstName
	userToUpdate.LastName = user.LastName
	if err := r.storage.Save(&userToUpdate).Error; err != nil {
		return err
	}
	return nil

}

func (r *userRepo) AddUserToRole(id string, role models.Role) error {
	userToUpdate := models.User{}
	if err := r.storage.Find(&userToUpdate).Where("id =?", id).Error; err != nil {
		return err
	}
	userToUpdate.Roles = append(userToUpdate.Roles, role)
	if err := r.storage.Save(&userToUpdate).Error; err != nil {
		return err
	}
	return nil
}

// get user by id
func (r *userRepo) GetById(id string) (models.User, error) {
	user := models.User{}
	if err := r.storage.Find(&user).Where("id =?", id).Error; err != nil {
		return user, err
	}
	return user, nil
}

// Activate user by id
func (r *userRepo) Activate(id string) error {
	user := models.User{}
	if err := r.storage.Find(&user).Where("id =?", id).Error; err != nil {
		return err
	}
	user.ActivatedAt = sql.NullTime{Time: time.Now()}
	user.Active = true
	if err := r.storage.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

// Deactivate user by id
func (r *userRepo) Deactivate(id string) error {
	user := models.User{}
	if err := r.storage.Find(&user).Where("id =?", id).Error; err != nil {
		return err
	}
	user.ActivatedAt = sql.NullTime{}
	user.Active = false
	if err := r.storage.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepo) GetByUsername(username string) (models.User, error) {
	user := models.User{}
	if err := r.storage.Find(&user).Where("username =?", username).Error; err != nil {
		return user, err
	}
	return user, nil
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
