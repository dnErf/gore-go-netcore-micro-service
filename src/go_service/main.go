package main

import (
	"encoding/json"
	"fmt"
	"gore/views"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", views.RenderHome).Methods("GET")
	router.HandleFunc("/test-service", HandleTestService).Methods("POST")

	router.PathPrefix("/bin/").Handler(http.StripPrefix("/bin/", http.FileServer(http.Dir("bin"))))

	fmt.Print("listening to :3000\n")
	http.ListenAndServe(":3000", router)
}

type DataObject struct {
	Message string `json:"message"`
}

type RequestPayload struct {
	Action string     `json:"action"`
	Data   DataObject `json:"data"`
}

type ResponsePaystub struct {
	IsFine  bool   `json:"is_fine"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func HandleTestService(w http.ResponseWriter, r *http.Request) {
	var payload RequestPayload
	err := ReadJSON(w, r, &payload)
	if err != nil {
		return
	}
	fmt.Printf("%s", payload)

	var paystub ResponsePaystub
	paystub.IsFine = true
	paystub.Message = "test is successful"
	WriteJSON(w, paystub)
}

const jsonMaxBytes int = 1024 * 10

func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	r.Body = http.MaxBytesReader(w, r.Body, int64(jsonMaxBytes))
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return err
	}
	return nil
}

func WriteJSON(w http.ResponseWriter, data any) error {
	output, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(output)
	if err != nil {
		return err
	}

	return nil
}
