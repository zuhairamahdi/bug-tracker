package core

import (
	"bugtracker/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).UTC()
	claims["username"] = user.Username
	secretbox := []byte(os.Getenv("JWT_SECRET"))
	if len(secretbox) == 0 {
		secretbox = []byte("secret")
	}
	tokenString, err := token.SignedString(secretbox)
	return tokenString, err
}
