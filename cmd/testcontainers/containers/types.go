package containers

import (
	"context"
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/internal/infrastructure/repository/userrepository"
	"github.com/bozd4g/fb.testcontainers/pkg/rabbitmq"
	"github.com/ory/dockertest"
	"gorm.io/gorm"
)

type TestContainer struct {
	Ctx context.Context

	Repository RepositoryContainer
	RabbitMq   RabbitMqContainer

	pool      dockertest.Pool
	resources []dockertest.Resource
}
type ITestContainer interface{}

type RepositoryContainer struct {
	Db             gorm.DB
	UserRepository userrepository.IUserRepository
}

type RabbitMqContainer struct {
	Opts   rabbitmq.Opts
	Broker rabbitmq.IRabbitMq
}
