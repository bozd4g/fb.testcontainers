package userservice

import "github.com/bozd4g/fb.testcontainers/pkg/rabbitmq"

type IUserService interface{}

type UserService struct {
	rabbitmq rabbitmq.IRabbitMq
}
