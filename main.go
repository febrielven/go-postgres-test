package main

import (
	"fmt"
	hd "go-postgres-test/apiV2/handlers"
	"go-postgres-test/config"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	// call connection db
	db := config.CreateConnection()
	defer db.Close()
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	// Create repos
	rHandler := hd.NewRideHandler(db)
	r.Route("/api/", func(rt chi.Router) {
		rt.Mount("/rides", rideRouter(rHandler))
	})

	fmt.Println("Starting server on the port 8082...")
	log.Fatal(http.ListenAndServe(":8082", r))
}

// A completely separate router for rides routes
func rideRouter(rHandler *hd.RideHandler) http.Handler {
	r := chi.NewRouter()
	r.Get("/", rHandler.Fetch)
	r.Get("/{id}", rHandler.GetByID)
	r.Post("/", rHandler.Save)
	return r
}
