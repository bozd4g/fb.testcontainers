package userrepository

import (
	"errors"
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/internal/domain/user"
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/internal/infrastructure/brokerconsts"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func New(db *gorm.DB) IUserRepository {
	return UserRepository{db: db}
}

func (repository UserRepository) Add(entity user.Entity) (*user.CreatedEvent, error) {
	transaction := repository.db.Begin()
	defer func() (*user.CreatedEvent, error) {
		if r := recover(); r != nil {
			transaction.Rollback()
		}

		return nil, errors.New("An error occured while creating a new user!")
	}()

	entity.Id = uuid.New()
	transaction.Create(&entity)
	if err := transaction.Error; err != nil {
		transaction.Rollback()
		return nil, err
	}

	transaction.Commit()
	return &user.CreatedEvent{
		ExchangeName: brokerconsts.UserCreatedExchangeName,
		Id:           entity.Id,
	}, nil
}

func (repository UserRepository) Get(id uuid.UUID) (*user.Entity, error) {
	var entity user.Entity
	result := repository.db.First(&entity, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &entity, nil
}

func (repository UserRepository) GetAll() ([]user.Entity, error) {
	var users []user.Entity
	result := repository.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
