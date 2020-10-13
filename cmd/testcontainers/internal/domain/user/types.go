package user

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Surname  string    `json:"surname"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}