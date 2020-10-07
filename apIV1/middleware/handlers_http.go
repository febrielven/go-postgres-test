package middleware

import (
	"fmt"
	// package db sql
	"encoding/json" //package to encode or decode the json into struct  and vice versa
	"log"
	"net/http"
	"strconv" // package used to convert string to int type

	"github.com/febrielven/go-postgres-test/apiV1/models" // models package where ride schema is defined

	"github.com/gorilla/mux" // used to get params form routes
	_ "github.com/lib/pq"    // postgres golang driver
)

// Response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

//CreateRides will return all user
func CreateRides(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// create an empty ride of  type Models.Rides
	var ride models.Rides

	err := json.NewDecoder(r.Body).Decode(&ride)

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	// call insert user function and pass the user
	insertID := insertRide(ride)

	// format a response object
	res := response{
		ID:      insertID,
		Message: "Ride created successfully",
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// GetRide will return a single ride by its id
func GetRide(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// get rideid from the request params, key is "id"
	params := mux.Vars(r)

	// convert string id to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("unable to convert string to int, %v", err)
	}

	// call the GetRideByID function with ride id to retrieve a single ride
	ride, err := getRideByID(int64(id))

	if err != nil {
		log.Fatalf("unable to get ride , %v", err)
	}

	json.NewEncoder(w).Encode(ride)

}

// GetRides will return a lists rides
func GetRides(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// call the getAllRides function with return lists rides
	rides, err := getAllRides()

	if err != nil {
		log.Fatalf("unable to get ride, %v", err)
	}
	json.NewEncoder(w).Encode(rides)
}

// UpdateRide detail in the postgres db
func UpdateRide(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the userid from the request params, key is id
	params := mux.Vars(r)

	// convert string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string to int. %v", err)
	}
	// create an empty rides of type models.rides
	var ride models.Rides

	// decode the json request to ride
	err = json.NewDecoder(r.Body).Decode(&ride)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call update ride to update the ride
	updateRows := updateRide(int64(id), ride)
	// format the message string
	msg := fmt.Sprintf("Ride updated successfully, Total rows/record affected %v", updateRows)

	// format the json message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)

}

// DeleteRide detail in the postgres db
func DeleteRide(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get rideid from the request params, key is "id"
	params := mux.Vars(r)

	// convert id string to integer
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable convert the string to integer. %v", err)
	}

	// call the deleteRide, convert the int to int64
	deleteRows := deleteRide(int64(id))

	// format message string
	msg := fmt.Sprintf("Ride delete successfully, Total rows/record affected. %v", deleteRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}
	// send the response
	json.NewEncoder(w).Encode(res)
}
