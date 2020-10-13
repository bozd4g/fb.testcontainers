package rabbitmq

import (
	"log"
	"os"
)

func (broker *RabbitMq) Consume(queueName string, prefetchCount int, onConsumed func(message []byte)) error {
	err := broker.channel.Qos(prefetchCount, 0, false)
	if err != nil {
		return err
	}

	consumerChannel, err := broker.channel.Consume(queueName,"",false,false,false,false,nil)
	if err != nil {
		return err
	}

	stopChan := make(chan bool)

	go func() {
		log.Printf("Consumer is ready, PID: %d", os.Getpid())
		for d := range consumerChannel {
			onConsumed(d.Body)

			if err := d.Ack(false); err != nil {
				log.Printf("Error acknowledging message: %s", err)
			}
		}
	}()

	<-stopChan

	return nil
}