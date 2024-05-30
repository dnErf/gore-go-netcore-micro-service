package main

import (
	"bff/web"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", web.RenderHome).Methods("GET")

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer((http.Dir("assets")))))

	fmt.Print("listening to :9999\n")
	http.ListenAndServe(":9999", router)
}
