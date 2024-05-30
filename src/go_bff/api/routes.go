package api

import (
	"bff/internal/models"
	"bff/internal/utils"
	"fmt"
	"net/http"
)

const TestServiceRoute = ""

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
}
