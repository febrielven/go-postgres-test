package repository

import (
	"context"
	"database/sql"
	models "go-postgres-test/apiV2/models"
)

// RideRepository ...
type RideRepository interface {
	// Fetch return lists rides, message and error
	Fetch(ctx context.Context) ([]*models.Rides, error)
	// // GeByID return rows rides, message and error
	// GeByID(ctx context.Context, id int64) (*models.Rides, error)
	// // Create return id, message and error
	// Create(ctx context.Context, ride *models.Rides) (int64, error)
	// // Update return id, message and error
	// Update(ctx context.Context, ride *models.Rides) (int64, error)
	// // Delete return id, message and error
	// Delete(ctx context.Context, id int64) (bool, error)
}

// RideRepo implements RideRepository
type RideRepo struct {
	db *sql.DB
}

// NewRideRepo ...
func NewRideRepo(db *sql.DB) *RideRepo {
	return &RideRepo{
		db: db,
	}
}

// Fetch ...
func (p *RideRepo) Fetch(ctx context.Context) ([]*models.Rides, error) {
	// create the sql statement
	sqlStatement := "SELECT * FROM rides"
	return p.fetch(ctx, sqlStatement)
}

// get from the DB
func (p *RideRepo) fetch(ctx context.Context, sqlStatement string) ([]*models.Rides, error) {

	// execute query statement
	rows, err := p.db.QueryContext(ctx, sqlStatement)

	if err != nil {
		// log.Fatalf("Unable to execute the query. %v", err)
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
			// log.Fatalf("Unable to scan the row. %v", err)
			return nil, err
		}
		// append the data in the rides slice
		payload = append(payload, data)
	}

	return payload, nil

}
