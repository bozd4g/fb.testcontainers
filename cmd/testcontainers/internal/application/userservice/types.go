package userservice

import (
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/internal/infrastructure/repository/userrepository"
	"github.com/bozd4g/fb.testcontainers/pkg/rabbitmq"
	"github.com/google/uuid"
)

type IUserService interface {
	Create(entity UserCreateRequestDto) error
	GetAll() ([]UserDto, error)
	Get(id uuid.UUID) (*UserDto, error)
}

type UserService struct {
	broker     rabbitmq.IRabbitMq
	repository userrepository.IUserRepository
}
