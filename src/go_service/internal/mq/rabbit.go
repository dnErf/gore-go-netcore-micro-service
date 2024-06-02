package mq

import (
	"encoding/json"
	"fmt"
	"gore/internal/models"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitClientSession struct {
	AmqpConnection *amqp.Connection
	AmqpChannel    *amqp.Channel
	AmqpQueue      amqp.Queue
}

var RabbitSession RabbitClientSession

func StartSessionToRabbitMq(c *amqp.Connection) {
	// TODO review the use of context
	// ctx, dn := context.WithCancel(context.Background())
	// sessions := make(chan RabbitClientSession)

	go func() {
		ch, err := c.Channel()
		if err != nil {
			log.Fatalf("cannot create channel %v", err)
		}

		if err := ch.ExchangeDeclare("topic_exchange", "topic", false, true, false, false, nil); err != nil {
			log.Fatalf("cannot declare topic exchange %v", err)
		}

		q, err := ch.QueueDeclare("", false, false, true, false, nil)
		if err != nil {
			log.Fatalf("cannot declare queue %v", err)
		}

		RabbitSession := RabbitClientSession{c, ch, q}

		messages, err := RabbitSession.BindQueueChannelToTopics([]string{"log.INFO"})
		if err != nil {
			log.Fatalf("cannot get delivery messages %v", err)
		}

		for d := range messages {
			var payload models.ResponsePaystub
			_ = json.Unmarshal(d.Body, &payload)
			fmt.Printf("%v", payload)
		}
	}()
}

func (rcs *RabbitClientSession) BindQueueChannelToTopics(topics []string) (<-chan amqp.Delivery, error) {
	for _, s := range topics {
		err := rcs.AmqpChannel.QueueBind(rcs.AmqpQueue.Name, s, "topic_exchange", false, nil)

		if err != nil {
			log.Fatalf("cannot bind topic to queue %v", err)
		}
	}

	return rcs.AmqpChannel.Consume(rcs.AmqpQueue.Name, "", true, false, false, false, nil)
}
