package api

import (
	"fmt"
	"gore/internal/models"
	"gore/internal/utils"
	"net/http"
)

func HandleTestService(w http.ResponseWriter, r *http.Request) {
	var payload models.RequestPayload
	err := utils.ReadJSON(w, r, &payload)
	if err != nil {
		return
	}
	fmt.Printf("%s", payload)

	var paystub models.ResponsePaystub
	paystub.IsFine = true
	paystub.Message = "test is successful"
	utils.WriteJSON(w, paystub)
}
