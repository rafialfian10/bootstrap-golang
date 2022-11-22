package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter() // buat router dengan mux
	port := "5000"

	// Routing
	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}).Methods("GET")

	route.HandleFunc("about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About"))
	}).Methods("GET")

	fmt.Println("Server sedang berjalan di port " + port)
	http.ListenAndServe("Localhost:"+port, route)
}
