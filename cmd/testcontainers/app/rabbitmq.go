package app

import (
	"fmt"
	"github.com/bozd4g/fb.testcontainers/cmd/testcontainers/internal/infrastructure/brokerconsts"
	"github.com/bozd4g/fb.testcontainers/pkg/rabbitmq"
)

func (application *Application) AddRabbitMq(opts rabbitmq.Opts) *Application {
	broker, err := rabbitmq.New(fmt.Sprintf("amqp://%s:%s@%s/%s", opts.Username, opts.Password, opts.Host, opts.VirtualHost))
	if err != nil {
		application.logger.Error("An error occured while connection to rabbitmq! ", err)
		return application
	}

	application.broker = broker
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