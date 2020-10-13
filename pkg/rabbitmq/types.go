package rabbitmq

import "github.com/streadway/amqp"

type IRabbitMq interface {
	DeclareQueue()
	Publish(queueName string, body string)
}

type RabbitMq struct {
	channel *amqp.Channel
}
