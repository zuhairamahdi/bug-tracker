package structs

import "bugtracker/models"

type NewUser struct {
	Username  string        `json:"username"`
	Email     string        `json:"email"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Password  string        `json:"password"`
	Roles     []models.Role `json:"roles"`
}

type User struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Active    bool   `json:"active"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserWithRoles struct {
	User
	Roles []Role `json:"roles"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
