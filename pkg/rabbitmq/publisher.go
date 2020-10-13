package rabbitmq

import (
	"github.com/streadway/amqp"
)

func (broker *RabbitMq) Publish(exchangeName string, body []byte) error {
	return broker.channel.Publish(exchangeName, "", false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
}
