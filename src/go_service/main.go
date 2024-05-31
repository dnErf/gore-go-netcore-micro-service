package main

import (
	"fmt"
	"gore/api"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/test-service", api.HandleTestService).Methods("POST")

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	fmt.Print("listening to :9998\n")
	http.ListenAndServe(":9998", router)
}
