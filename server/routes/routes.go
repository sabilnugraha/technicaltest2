package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	Token(r)
	Data(r)
}
