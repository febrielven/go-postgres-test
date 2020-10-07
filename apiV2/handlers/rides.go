package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv" // package used to convert string to int type

	"github.com/febrielven/go-postgres-test/apiV2/models"
	repository "github.com/febrielven/go-postgres-test/apiV2/repository"
	rides "github.com/febrielven/go-postgres-test/apiV2/repository/rides"

	"github.com/go-chi/chi" // used to get params form routes
)

// RideHandler will hold everything that controller needs
type RideHandler struct {
	rideRepo repository.RideRepository
}

// NewRideHandler returns a new RideHandler
func NewRideHandler(db *sql.DB) *RideHandler {
	return &RideHandler{
		rideRepo: rides.NewRideRepo(db),
	}
}

// Fetch ...
func (ride *RideHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	//
	payload, err := ride.rideRepo.Fetch(r.Context())

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unknown error")
		return
	}
	respondWithJSON(w, http.StatusOK, payload)
}

// GetByID ...
func (ride *RideHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get rideid from request params, key is "id"
	params := chi.URLParam(r, "id")

	// convert string to int
	id, err := strconv.Atoi(params)
	if err != nil {
		fmt.Printf("unable to convert string to int %v", err)
		respondWithError(w, http.StatusUnprocessableEntity, "unable to convert string to int")
		return
	}
	payload, err := ride.rideRepo.GetByID(r.Context(), int64(id))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unknow error")
		return

	}
	if payload == nil && err == nil {
		respondWithError(w, http.StatusUnprocessableEntity, "Could not find any rides")
		return
	}

	respondWithJSON(w, http.StatusOK, payload)

}

// Save ...
func (ride *RideHandler) Save(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// create empty ride of models.Rides
	var mride models.Rides

	err := json.NewDecoder(r.Body).Decode(&mride)

	if err != nil {
		fmt.Printf("unable to decoder the request body, %v", err)
		respondWithError(w, http.StatusUnprocessableEntity, "unable to decoder the request body")
		return
	}

	// call save ride repository and pass the ride
	insertedID, err := ride.rideRepo.Save(r.Context(), mride)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unknow error")
		return
	}
	// format the json message
	res := response{
		ID:      insertedID,
		Message: "Ride created successfully",
	}
	// send the response
	respondWithJSON(w, http.StatusCreated, res)
}

// Update ..
func (ride *RideHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// get rideid from request params, key is "id"
	params := chi.URLParam(r, "id")

	// convert string to int
	id, err := strconv.Atoi(params)
	if err != nil {
		fmt.Printf("unable to convert string to int %v", err)
		respondWithError(w, http.StatusUnprocessableEntity, "unable to convert string to int")
		return
	}

	// create empty ride of models.Rides
	var mride models.Rides
	mride.ID = int64(id)
	if err != nil {
		fmt.Printf("unable to decoder the request body, %v", err)
		respondWithError(w, http.StatusUnprocessableEntity, "unable to decoder the request body")
		return
	}

	// call update ride repository and pass the ride
	updateRows, err := ride.rideRepo.Update(r.Context(), mride)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "unknow error")
		return
	}

	msg := fmt.Sprintf("Ride updated successfully, Total rows/record affected %v", updateRows)
	// format the json message
	res := response{
		ID:      int64(id),
		Message: msg,
	}
	// send the response
	respondWithJSON(w, http.StatusOK, res)

}

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// respondWithJSON write json response format
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"message": msg})
}
