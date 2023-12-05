package appcontext

import (
	"bugtracker/structs"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func GetUserData(c *fiber.Ctx) (structs.User, error) {
	user := structs.User{}
	// convert(c.Locals("user"), &user)
	userJson, err := json.Marshal(c.Locals("user"))
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(userJson, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}
