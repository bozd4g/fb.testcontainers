package userservice

import "github.com/google/uuid"

type UserDto struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Surname  string    `json:"surname"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type UserCreateRequestDto struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
