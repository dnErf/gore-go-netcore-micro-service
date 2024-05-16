package main

import (
	"fmt"
	"gore/views"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", views.RenderHome).Methods("GET")

	fmt.Print("listening to :3000")
	http.ListenAndServe(":3000", router)
}
