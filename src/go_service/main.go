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

	router.PathPrefix("/bin/").Handler(http.StripPrefix("/bin/", http.FileServer(http.Dir("bin"))))

	fmt.Print("listening to :3000\n")
	http.ListenAndServe(":3000", router)
}
