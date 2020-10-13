package rabbitmq

import "github.com/streadway/amqp"

type IRabbitMq interface {
	DeclareQueue(name string) (amqp.Queue, error)
	Publish(exchangeName string, body []byte) error
	Consume(queueName string, prefetchCount int, onConsumed func(message []byte)) error
}

type RabbitMq struct {
	channel *amqp.Channel
}
