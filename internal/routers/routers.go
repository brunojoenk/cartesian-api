package routers

import (
	"github.com/gorilla/mux"

	handlers "github.com/brunojoenk/cartesian-api/internal/handlers"
)

func CreateRouters() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/points", handlers.HandleDistance).Methods("GET")
	return router
}
