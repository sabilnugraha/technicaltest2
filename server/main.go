package main

import (
	"fmt"
	"net/http"

	"technicaltest/database"
	"technicaltest/pkg/mysql"
	"technicaltest/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	mysql.DatabaseInit()
	database.RunMigration()

	r := mux.NewRouter()
	routes.RouteInit(r)

	var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Content-Type", "Authorization"})
	var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})

	var port = "5000"
	fmt.Println("server running localhost:" + port)
	http.ListenAndServe("localhost:"+port, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))
}
