package handlers

import (
	"encoding/json"
	repository "go-postgres-test/apiV2/repository"
	"net/http"
)

// RideHandler will hold everything that controller needs
type RideHandler struct {
	rideRepo repository.RideRepository
}

// NewRideHandler returns a new RideHandler
func NewRideHandler(rideRepo repository.RideRepository) *RideHandler {
	return &RideHandler{
		rideRepo: rideRepo,
	}
}

// Fetch ...
func (ride *RideHandler) Fetch(w http.ResponseWriter, r *http.Request) {

	payload, _ := ride.rideRepo.Fetch(r.Context())

	respondwithJSON(w, http.StatusOK, payload)
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
