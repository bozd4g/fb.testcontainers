package userservice

import (
	"fmt"
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/internal/domain/user"
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/internal/infrastructure/repository/userrepository"
	"github.com/bozd4g/fb.testcontainers/pkg/rabbitmq"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

func New(broker rabbitmq.IRabbitMq, repository userrepository.IUserRepository) IUserService {
	return UserService{rabbitmq: broker, repository: repository}
}

func (service UserService) Create(userDto UserDto) error {
	var entity user.Entity
	err := mapstructure.Decode(userDto, &entity)
	if err != nil {
		return err
	}

	event, err := service.repository.Add(entity)
	if err != nil {
		return err
	}

	// TODO: Throw event to queue
	fmt.Println(event)
	return nil
}

func (service UserService) GetAll() ([]UserDto, error) {
	users, err := service.repository.GetAll()
	if err != nil {
		return nil, err
	}

	var dtos []UserDto
	err = mapstructure.Decode(users, &dtos)
	if err != nil {
		return nil, err
	}

	return dtos, nil
}

func (service UserService) Get(id uuid.UUID) (*UserDto, error) {
	user, err := service.repository.Get(id)
	if err != nil {
		return nil, err
	}

	var dto UserDto
	err = mapstructure.Decode(user, &dto)
	if err != nil {
		return nil, err
	}

	return &dto, nil
}
