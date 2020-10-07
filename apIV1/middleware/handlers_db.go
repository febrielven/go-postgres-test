package middleware

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/febrielven/go-postgres-test/apiV1/models"

	"github.com/joho/godotenv"
)

// createConnection with postgress
func createConnection() *sql.DB {
	//load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	//Cek Connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully Connection")

	//Return the connection
	return db
}

// insertRide one ride in the DB
func insertRide(ride models.Rides) int64 {
	// create connection in db postgres
	db := createConnection()

	// clone db connection
	defer db.Close()

	// create the insert sql query
	// returning rideid will return the id of inserted ride
	sqlStatement := `
		INSERT INTO rides (
			startlat,
			startlong,
			endlat,
			endlong,
			ridername,
			drivername,
			drivervehicle
		)VALUES(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7
		)RETURNING rideid`
	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// scan function will save the insert id in then id
	err := db.QueryRow(
		sqlStatement,
		ride.StartLat,
		ride.StartLong,
		ride.EndLat,
		ride.EndLong,
		ride.RiderName,
		ride.DriverName,
		ride.DriverVehicle).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}

	fmt.Printf("Inserted a one record %v", id)

	return id

}

// getRideByID ride from the db by rideid
// params ride id
func getRideByID(id int64) (models.Rides, error) {

	// create connection in db postgress
	db := createConnection()

	// clone db connection
	defer db.Close()

	// create a ride of models.Ride type
	var ride models.Rides

	// create the select sql query
	sqlStatement := `SELECT * FROM rides where rideid=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// unmarshal the row object the rides
	err := row.Scan(
		&ride.ID,
		&ride.StartLat,
		&ride.StartLong,
		&ride.EndLat,
		&ride.EndLong,
		&ride.RiderName,
		&ride.DriverName,
		&ride.DriverVehicle,
		&ride.Created)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
	case nil:
		return ride, nil
	default:
		log.Fatalf("unable to scan the row .%v", err)

	}

	return ride, err
}

// getAllRides from the DB
func getAllRides() ([]models.Rides, error) {

	// create the postgres db connection
	db := createConnection()

	// close the db connect
	defer db.Close()

	// create rides of models.Rides
	var rides []models.Rides

	// create the sql statement
	sqlStatement := `SELECT * FROM rides`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var ride models.Rides

		// unmarshal the row object to ride
		err := rows.Scan(
			&ride.ID,
			&ride.StartLat,
			&ride.StartLong,
			&ride.EndLat,
			&ride.EndLong,
			&ride.RiderName,
			&ride.DriverName,
			&ride.DriverVehicle,
			&ride.Created)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the rides slice
		rides = append(rides, ride)
	}

	return rides, err

}

// UpdateRide form db return int rideid
func updateRide(id int64, ride models.Rides) int64 {
	// create connection with db posgress
	db := createConnection()

	// close db connect
	defer db.Close()

	// create statement query
	sqlStatement := `UPDATE
				rides
			SET
				startlat=$2,
				startlong=$3,
				endlat= $4,
				endlong= $5,
				ridername=$6,
				drivername=$7,
				drivervehicle=$8
			WHERE
				rideid = $1`
	// excute query update
	res, err := db.Exec(
		sqlStatement,
		ride.ID,
		ride.StartLat,
		ride.StartLong,
		ride.EndLat,
		ride.EndLong,
		ride.RiderName,
		ride.DriverName,
		ride.DriverVehicle)

	if err != nil {
		log.Fatalf("Unable execute query update. %v", err)
	}

	rowAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/recourd affected %v", rowAffected)
	return rowAffected

}

// DeleteRide in the DB
func deleteRide(id int64) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM rides WHERE rideid=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}
