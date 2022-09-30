package routes

import (
	"technicaltest/handlers"

	"github.com/gorilla/mux"
)

func Token(r *mux.Router) {

	r.HandleFunc("/token", handlers.GetToken).Methods("GET")
}
