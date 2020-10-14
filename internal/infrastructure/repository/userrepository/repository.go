package userrepository

import (
	"github.com/bozd4g/fb.testcontainers/internal/domain/user"
	"github.com/google/uuid"
	"github.com/hashicorp/go-memdb"
)

func New() (IUserRepository, error) {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"users": &memdb.TableSchema{
				Name: "users",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Id"},
					},
					"name": &memdb.IndexSchema{
						Name:    "name",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
					"surname": &memdb.IndexSchema{
						Name:    "surname",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Surname"},
					},
					"email": &memdb.IndexSchema{
						Name:    "email",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Email"},
					},
					"password": &memdb.IndexSchema{
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

	return UserRepository{db: db}, nil
}

func (repository UserRepository) Add(entity user.Entity) user.CreatedEvent {
	return user.CreatedEvent{}
}

func (repository UserRepository) Find(id uuid.UUID) {}
