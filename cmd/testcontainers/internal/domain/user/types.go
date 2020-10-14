package user

import "github.com/google/uuid"

type Entity struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Surname  string    `json:"surname"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type CreatedEvent struct {
	Id uuid.UUID `json:"id"`
}