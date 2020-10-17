package app

import (
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/internal/infrastructure/brokerconsts"
	"github.com/bozd4g/fb.testcontainers/pkg/rabbitmq"
)

func (application *Application) AddRabbitMq() *Application {
	broker, err := rabbitmq.New("amqp://guest:123456@localhost/demand")
	if err != nil {
		application.logger.Errorf("An error occured while connection to rabbitmq! Error: %+v", err)
		return application
	}

	application.broker = broker
	defer broker.Close()

	application.InitUserCreatedEvent()
	return application
}

func (application *Application) InitUserCreatedEvent() {
	_, err := application.broker.DeclareQueue(brokerconsts.UserCreatedQueueName)
	if err != nil {
		panic(err)
	}
	err = application.broker.DeclareExhange(brokerconsts.UserCreatedExchangeName)
	if err != nil {
		panic(err)
	}

	err = application.broker.BindQueue(brokerconsts.UserCreatedQueueName, brokerconsts.UserCreatedExchangeName)
	if err != nil {
		panic(err)
	}
}