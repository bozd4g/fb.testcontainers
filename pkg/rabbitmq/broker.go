package rabbitmq

import "github.com/streadway/amqp"

func New(ampqUrl string) (IRabbitMq, error) {
	connection, err := amqp.Dial(ampqUrl)
	if err != nil {
		return nil, err
	}

	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	defer channel.Close()

	return &RabbitMq{channel: channel}, nil
}

func (broker *RabbitMq) DeclareQueue(name string) (amqp.Queue, error) {
	return broker.channel.QueueDeclare(name, true, false, false, false, nil)
}