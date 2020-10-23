package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Entity struct {
	gorm.Model
	Id        uuid.UUID `gorm:"primaryKey, column:id" json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	IsDeleted bool      `gorm:"column:isDeleted" json:"isDeleted"`
}

type CreatedEvent struct {
	ExchangeName string    `json:"-"`
	Id           uuid.UUID `json:"id"`
}
