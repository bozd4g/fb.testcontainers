package userrepository

import (
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/internal/domain/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Add(entity user.Entity) (*user.CreatedEvent, error)
	Get(id uuid.UUID) (*user.Entity, error)
	GetAll() ([]user.Entity, error)
}

type UserRepository struct {
	db    *gorm.DB
}
