package rides

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	models "github.com/febrielven/go-postgres-test/apiV2/models"
	repository "github.com/febrielven/go-postgres-test/apiV2/repository"
)

// RideRepo ...
type RideRepo struct {
	db *sql.DB
}

// NewRideRepo implements repository.RideRepository ...
func NewRideRepo(db *sql.DB) repository.RideRepository {
	return &RideRepo{
		db: db,
	}
}

// Fetch ...
func (pg *RideRepo) Fetch(ctx context.Context) ([]*models.Rides, error) {
	// create the sql statement
	sqlStatement := "SELECT * FROM rides"
	return pg.fetch(ctx, sqlStatement)
}

// GetByID ...
func (pg *RideRepo) GetByID(ctx context.Context, id int64) (*models.Rides, error) {
	sqlStatement := "SELECT * FROM rides where rideid= $1"

	// create a ride of models.Ride type
	var ride models.Rides

	// execute sql statement
	rows := pg.db.QueryRow(sqlStatement, id)

	// unmarshal the row object the rides
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

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
		return nil, nil
	case nil:
		return &ride, nil
	default:
		fmt.Printf("unable to scan the row %v", err)
		return nil, err
	}

}

// Save ..
func (pg *RideRepo) Save(ctx context.Context, ride models.Rides) (int64, error) {

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

	// the inserted id will store this rideid
	var rideid int64

	// execute query sqlStatement
	// scan function will save the insert ride in then rideid
	err := pg.db.QueryRow(
		sqlStatement,
		ride.StartLat,
		ride.StartLong,
		ride.EndLat,
		ride.EndLong,
		ride.RiderName,
		ride.DriverName,
		ride.DriverVehicle).Scan(&rideid)

	if err != nil {
		fmt.Printf("unable to execute the query %v", err)
		return rideid, err
	}

	return rideid, nil
}

// Update ..
func (pg *RideRepo) Update(ctx context.Context, ride models.Rides) (int64, error) {

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

	// the inserted id will store this rideid
	var rideid int64

	// execute query sqlStatement
	// scan function will updated the insert ride in then rideid
	res, err := pg.db.Exec(
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
		fmt.Printf("unable to execute the query %v", err)
		return rideid, err
	}
	rowAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/recourd affected %v", rowAffected)
	return rowAffected, nil

}

// get from the DB
func (pg *RideRepo) fetch(ctx context.Context, sqlStatement string) ([]*models.Rides, error) {

	// execute query statement
	rows, err := pg.db.QueryContext(ctx, sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
		return nil, err
	}
	// close the statement
	defer rows.Close()

	// create payload of models.Rides
	payload := make([]*models.Rides, 0)

	// iterate  over the rows
	for rows.Next() {
		data := new(models.Rides)

		// unmarshal the row object to ride
		err := rows.Scan(
			&data.ID,
			&data.StartLat,
			&data.StartLong,
			&data.EndLat,
			&data.EndLong,
			&data.RiderName,
			&data.DriverName,
			&data.DriverVehicle,
			&data.Created)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
			return nil, err
		}
		// append the data in the rides slice
		payload = append(payload, data)
	}

	return payload, nil

}
