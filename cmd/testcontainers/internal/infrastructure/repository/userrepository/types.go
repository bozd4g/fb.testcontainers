package userrepository

import (
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/internal/domain/user"
	"github.com/google/uuid"
	"github.com/hashicorp/go-memdb"
)

type IUserRepository interface {
	Add(entity user.Entity) (*user.CreatedEvent, error)
	Get(id uuid.UUID) (*user.Entity, error)
	GetAll() ([]user.Entity, error)
}

type UserRepository struct {
	table string
	db    *memdb.MemDB
}
