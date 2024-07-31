package main

import (
	"api/config"
	"api/router"
	"api/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()
	container, err := services.NewContainer(cfg)
	if err != nil {
		log.Fatalf("Error creating container: %v", err)
	}
	Router := mux.NewRouter().StrictSlash(true)
	Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API INICIADA")
	}).Methods("GET")
	Router.PathPrefix("/predios").Handler(router.RouterPredio(container.PredioRepo))
	log.Fatal(http.ListenAndServe(":3000", Router))
}
