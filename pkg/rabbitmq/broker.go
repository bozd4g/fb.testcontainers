package rabbitmq

import "github.com/streadway/amqp"

func New(ampqUrl string) (IRabbitMq, error) {
	connection, err := amqp.Dial(ampqUrl)
	if err != nil {
		return nil, err
	}

	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMq{channel: channel}, nil
}

func (broker *RabbitMq) DeclareExhange(name string) error {
	return broker.channel.ExchangeDeclare(name, "fanout", true, false, false, false, nil)
}

func (broker *RabbitMq) DeclareQueue(name string) (amqp.Queue, error) {
	return broker.channel.QueueDeclare(name, true, false, false, false, nil)
}

func (broker *RabbitMq) BindQueue(queueName string, exchangeName string) error {
	return broker.channel.QueueBind(queueName, "", exchangeName, false, nil)
}

func (broker *RabbitMq) Close() {
	broker.connection.Close()
	broker.channel.Close()
}
