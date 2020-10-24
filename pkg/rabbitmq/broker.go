package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

func New(opts Opts) (IRabbitMq, error) {
	amqpString := fmt.Sprintf("amqp://%s:%s@%s/%s", opts.Username, opts.Password, opts.Host, opts.VirtualHost)
	connection, err := amqp.Dial(amqpString)
	if err != nil {
		return nil, err
	}

	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMq{connection: connection, channel: channel}, nil
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
