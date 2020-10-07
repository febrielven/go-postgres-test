package routes

import (
	"github.com/febrielven/go-postgres-test/apiV1/middleware"

	"github.com/gorilla/mux"
)

// Router is exported in main.go
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/rides", middleware.CreateRides).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/rides/{id}", middleware.GetRide).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/rides", middleware.GetRides).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/rides/{id}", middleware.UpdateRide).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/delete-ride/{id}", middleware.DeleteRide).Methods("DELETE", "OPTION")

	return router
}
