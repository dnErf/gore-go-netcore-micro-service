package mq

import (
	"bff/internal/models"
	"bff/internal/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitClientConnection struct {
	AmqpConnection *amqp.Connection
}

func Init(c *amqp.Connection) *RabbitClientConnection {
	ch, err := c.Channel()
	if err != nil {
		log.Fatalf("cannot create channel %v", err)
	}
	defer ch.Close()

	if err := ch.ExchangeDeclare("topic_exchange", "topic", false, true, false, false, nil); err != nil {
		log.Fatalf("cannot declare topic exchange %v", err)
	}

	return &RabbitClientConnection{AmqpConnection: c}
}

func (rcs *RabbitClientConnection) emitEvent(event_key string, payload []byte) error {
	ch, err := rcs.AmqpConnection.Channel()
	if err != nil {
		log.Fatalf("cannot create channel %v", err)
	}
	defer ch.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err = ch.PublishWithContext(
		ctx,
		"topic_exchange",
		event_key,
		false,
		false,
		amqp.Publishing{
			ContentType:  "text/plain",
			DeliveryMode: amqp.Transient,
			Body:         payload,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (rcc *RabbitClientConnection) EmitLog(data any) error {
	payloadByte, err := utils.MarshalTabIndenter(data)
	if err != nil {
		log.Fatalf("marshal indent %v", err)
		return err
	}

	if err = rcc.emitEvent("log.INFO", payloadByte); err != nil {
		log.Fatalf("emit event %v", err)
		return err
	}

	return nil
}

func (rcc *RabbitClientConnection) TestEmitEvent() {
	var paystub models.ResponsePaystub
	paystub.IsFine = true
	paystub.Message = "test is successful"

	j, err := json.MarshalIndent(&paystub, "", "\t")
	if err != nil {
		log.Fatalf("marshal indent %v", err)
	}

	for i := 0; i <= 100; i++ {
		fmt.Printf("%v\n", i)
		if err = rcc.emitEvent("log.INFO", []byte(j)); err != nil {
			log.Fatalf("emit event %v", err)
		}
	}
}
