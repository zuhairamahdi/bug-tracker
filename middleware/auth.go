package middleware

import "github.com/gofiber/fiber/v2"

// JWTmiddleware is a middleware that will check if the user is authenticated and return the result if it is successful to next
func JWTmiddleware(c *fiber.Ctx) error {
	token := c.Get("token")
	if len(token) == 0 {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	return c.Next()
}
