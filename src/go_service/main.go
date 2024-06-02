package main

import (
	"fmt"
	"gore/api"
	"gore/internal/mq"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	c, err := amqp.Dial("amqp://guest:guest@172.29.248.117:5672")
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer c.Close()
	mq.StartSessionToRabbitMq(c)
	defer mq.RabbitSession.AmqpChannel.Close()
	defer mq.RabbitSession.AmqpConnection.Close()

	router := mux.NewRouter()
	router.HandleFunc("/test-service", api.HandleTestService).Methods("POST")

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	fmt.Print("listening to :9998\n")
	http.ListenAndServe(":9998", router)
}
