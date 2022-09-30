package routes

import (
	"technicaltest/handlers"

	"technicaltest/pkg/middleware"
	"technicaltest/pkg/mysql"
	"technicaltest/repositories"

	"github.com/gorilla/mux"
)

func Data(r *mux.Router) {
	dataRepository := repositories.RepositoryData(mysql.DB)
	h := handlers.HandleData(dataRepository)

	r.HandleFunc("/data", middleware.Auth(middleware.UploadFile(h.CreateData))).Methods("POST")
}
