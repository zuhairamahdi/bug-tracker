package middleware

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// JWTmiddleware is a middleware that will check if the user is authenticated and return the result if it is successful to next
func JWTmiddleware(c *fiber.Ctx) error {
	token := c.Request().Header.Peek("Authorization")
	if len(token) == 0 && !strings.HasPrefix(string(token), "Bearer ") {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	//Remove Bearer from token

	token_value := strings.Split(string(token), "Bearer ")[1]
	//decrypt the token
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(token_value, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	//Check if 'exp' is in the future using time.Now()
	exp, _ := time.Parse(time.RFC3339, fmt.Sprintf("%s", claims["exp"]))
	if exp.Before(time.Now()) {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	c.Locals("user", claims["user"])
	return c.Next()
}
