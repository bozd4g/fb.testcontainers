package app

import (
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/internal/infrastructure/brokerconsts"
	"github.com/bozd4g/fb.testcontainers/pkg/rabbitmq"
)

func (application *Application) AddRabbitMq(opts rabbitmq.Opts) *Application {
	broker, err := rabbitmq.New(opts)
	if err != nil {
		application.logger.Error("An error occured while connection to rabbitmq! ", err)
		return application
	}

	application.broker = broker
	application.InitUserCreatedEvent()

	return application
}

func (application *Application) InitUserCreatedEvent() {
	err := application.broker.Bind(brokerconsts.UserCreatedExchangeName, brokerconsts.UserCreatedQueueName)
	if err != nil {
		application.logger.Error("An error occured while binding to exchange! ", err)
	}
}