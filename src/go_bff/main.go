package main

import (
	"bff/api"
	"bff/internal/mq"
	"bff/web"
	"fmt"
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
	mq.StartSessionToRabbitMq(c)
	defer mq.RabbitSession.AmqpChannel.Close()
	defer mq.RabbitSession.AmqpConnection.Close()

	router := mux.NewRouter()
	router.HandleFunc(api.TestServiceRoute, api.TestServiceRouteHandler)
	router.HandleFunc("/", web.RenderHome).Methods("GET")

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer((http.Dir("assets")))))

	fmt.Print("listening to :9999\n")
	http.ListenAndServe(":9999", router)
}
