package userservice

import (
	"encoding/json"
	"fmt"
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/internal/domain/user"
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/internal/infrastructure/repository/userrepository"
	"github.com/bozd4g/fb.testcontainers/pkg/rabbitmq"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

func New(broker rabbitmq.IRabbitMq, repository userrepository.IUserRepository) IUserService {
	return UserService{broker: broker, repository: repository}
}

func (service UserService) Create(userDto UserCreateRequestDto) error {
	var entity user.Entity
	err := mapstructure.Decode(userDto, &entity)
	if err != nil {
		return err
	}

	event, err := service.repository.Add(entity)
	if err != nil {
		return err
	}

	jsonEvent, err := json.Marshal(&event)
	if err != nil {
		return err
	}

	err = service.broker.Publish(event.ExchangeName, jsonEvent)
	if err != nil {
		fmt.Println(fmt.Sprintf("An error occured while throwing the event! Event: %+v, Error: %v+", event, err))
		return err
	}
	return nil
}

func (service UserService) GetAll() ([]UserDto, error) {
	dtos := make([]UserDto, 0)

	users, err := service.repository.GetAll()
	if err != nil {
		return dtos, err
	}

	err = mapstructure.Decode(users, &dtos)
	if err != nil {
		return dtos, err
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
