package api

import (
	"bff/internal/models"
	"bff/internal/mq"
	"bff/internal/utils"
	"fmt"
	"net/http"
)

type Routing struct {
	rabbitmq *mq.RabbitClientConnection
}

func Init(rabbitmq *mq.RabbitClientConnection) *Routing {
	return &Routing{
		rabbitmq: rabbitmq,
	}
}

const TestServiceRoute = "/test-service"

func (options *Routing) TestServiceRouteHandler(w http.ResponseWriter, r *http.Request) {
	var payload models.RequestPayload
	err := utils.ReadJSON(w, r, &payload)
	if err != nil {
		return
	}
	fmt.Printf("%v", payload)

	var paystub models.ResponsePaystub
	paystub.IsFine = true
	paystub.Message = "test is successful"
	utils.WriteJSON(w, paystub)

	options.rabbitmq.EmitLog(paystub)
}

const LogRoute = "/log"

func (options *Routing) LogRouteHandler(w http.ResponseWriter, r *http.Request) {
	var payload models.RequestPayload
	err := utils.ReadJSON(w, r, &payload)
	if err != nil {
		return
	}
	fmt.Printf("%v", payload)

	options.rabbitmq.EmitLog(payload)
}
