package router

import (
	"api/handlers"
	"api/repository"

	"github.com/gorilla/mux"
)

func RouterPredio(repo repository.PredioRepository) *mux.Router {
	routerPredio := handlers.PredioHandler{Repo: repo}
	router := mux.NewRouter().StrictSlash(true).PathPrefix("/predios").Subrouter()
	router.HandleFunc("/", routerPredio.GetPredios).Methods("GET")
	router.HandleFunc("/", routerPredio.CreatePredio).Methods("POST")
	router.HandleFunc("/{t_id}", routerPredio.GetPredio).Methods("GET")
	router.HandleFunc("/{t_id}", routerPredio.DeletePredio).Methods("DELETE")
	router.HandleFunc("/{t_id}", routerPredio.UpdatePredio).Methods("PUT")
	return router
}
