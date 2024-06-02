package api

import (
	"bff/internal/models"
	"bff/internal/mq"
	"bff/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

const TestServiceRoute = "/test-service"

func TestServiceRouteHandler(w http.ResponseWriter, r *http.Request) {
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

	j, err := json.MarshalIndent(&payload, "", "\t")
	if err != nil {
		return
	}

	if err = mq.RabbitSession.EmitEvent("log.INFO", []byte(j)); err != nil {
		return
	}
}
