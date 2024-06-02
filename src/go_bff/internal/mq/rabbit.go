package mq

import (
	"bff/internal/models"
	"encoding/json"
	"fmt"
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

		RabbitSession.TestEmitEvent()
	}()
}

func (rcs *RabbitClientSession) EmitEvent(event_key string, payload []byte) error {
	err := rcs.AmqpChannel.Publish(
		"topic_exchange",
		event_key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        payload,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (rcs *RabbitClientSession) TestEmitEvent() {
	var paystub models.ResponsePaystub
	paystub.IsFine = true
	paystub.Message = "test is successful"

	j, err := json.MarshalIndent(&paystub, "", "\t")
	if err != nil {
		log.Fatalf("marshal indent %v", err)
	}

	for i := 0; i <= 100; i++ {
		fmt.Printf("%v\n", i)
		if err = rcs.EmitEvent("log.INFO", []byte(j)); err != nil {
			log.Fatalf("emit event %v", err)
		}
	}
}
