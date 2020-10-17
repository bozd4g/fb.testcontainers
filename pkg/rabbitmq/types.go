package rabbitmq

import "github.com/streadway/amqp"

type IRabbitMq interface {
	DeclareExhange(name string) error
	DeclareQueue(name string) (amqp.Queue, error)
	Publish(exchangeName string, body []byte) error
	BindQueue(queueName string, exchangeName string) error
	Consume(queueName string, prefetchCount int, onConsumed func(message []byte)) error
	Close()
}

type RabbitMq struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}
