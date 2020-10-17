package userrepository

import (
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/internal/domain/user"
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/internal/infrastructure/brokerconsts"
	"github.com/google/uuid"
	"github.com/hashicorp/go-memdb"
)

func New() (IUserRepository, error) {
	tableName := "users"
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			tableName: {
				Name: tableName,
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Id"},
					},
					"name": {
						Name:    "name",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
					"surname": {
						Name:    "surname",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Surname"},
					},
					"email": {
						Name:    "email",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Email"},
					},
					"password": {
						Name:    "password",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Password"},
					},
				},
			},
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		return nil, err
	}

	return UserRepository{db: db, table: tableName}, nil
}

func (repository UserRepository) Add(entity user.Entity) (*user.CreatedEvent, error) {
	transaction := repository.db.Txn(true)

	entity.Id = uuid.New()
	err := transaction.Insert(repository.table, entity)
	if err != nil {
		transaction.Abort()
		return nil, err
	}

	transaction.Commit()
	defer transaction.Abort()
	return &user.CreatedEvent{
		ExchangeName: brokerconsts.UserCreatedExchangeName,
		Id:           entity.Id,
	}, nil
}

func (repository UserRepository) Get(id uuid.UUID) (*user.Entity, error) {
	transaction := repository.db.Txn(false)
	entity, err := transaction.First(repository.table, "id", id)
	if err != nil {
		return nil, err
	}
	return entity.(*user.Entity), nil
}

func (repository UserRepository) GetAll() ([]user.Entity, error) {
	transaction := repository.db.Txn(false)
	iterator, err := transaction.Get(repository.table, "id")
	if err != nil {
		return nil, err
	}

	var users []user.Entity
	for obj := iterator.Next(); obj != nil; obj = iterator.Next() {
		p := obj.(user.Entity)
		users = append(users, p)
	}

	return users, nil
}
