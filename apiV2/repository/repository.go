package repository

import (
	"context"

	models "github.com/febrielven/go-postgres-test/apiV2/models"
)

// RideRepository ...
type RideRepository interface {
	// Fetch return lists rides, message and error
	Fetch(ctx context.Context) ([]*models.Rides, error)
	// GetByID return rows rides, message and error
	GetByID(ctx context.Context, id int64) (*models.Rides, error)
	// Save return id, message and error
	Save(ctx context.Context, ride models.Rides) (int64, error)
	// Update return id, message and error
	Update(ctx context.Context, ride models.Rides) (int64, error)
	// // Delete return id, message and error
	// Delete(ctx context.Context, id int64) (bool, error)
}
