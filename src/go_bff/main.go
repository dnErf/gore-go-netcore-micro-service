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
	defer c.Close()

	rabbitmq := mq.Init(c)
	routing := api.Init(rabbitmq)

	router := mux.NewRouter()
	router.HandleFunc(api.TestServiceRoute, routing.TestServiceRouteHandler)
	router.HandleFunc(api.LogRoute, routing.LogRouteHandler)
	router.HandleFunc("/", web.RenderHome).Methods("GET")

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer((http.Dir("assets")))))

	fmt.Print("listening to :9999\n")
	http.ListenAndServe(":9999", router)
}
