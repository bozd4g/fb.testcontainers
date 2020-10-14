package userrepository

import (
	"github.com/bozd4g/fb.testcontainers/internal/domain/user"
	"github.com/google/uuid"
	"github.com/hashicorp/go-memdb"
)

type IUserRepository interface {
	Add(entity user.Entity) user.CreatedEvent
	Find(id uuid.UUID)
}

type UserRepository struct {
	db *memdb.MemDB
}
