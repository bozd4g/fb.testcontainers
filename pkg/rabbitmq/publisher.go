package rabbitmq

import (
	"errors"
	"github.com/streadway/amqp"
)

func (broker *RabbitMq) Publish(exchangeName string, body []byte) error {
	if broker == nil || broker.channel == nil {
		return errors.New("Broker was not connected!")
	}
	return broker.channel.Publish(exchangeName, "", false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
}
