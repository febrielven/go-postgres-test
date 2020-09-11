package models

import (
	"time"
)

// Rides schema of rides table ...
type Rides struct {
	ID            int64     `sql:"AUTO_INCREMENT" json:"rideid"`
	StartLat      int16     `json:"startlat"`
	StartLog      int16     `json:"startlog"`
	EndLat        int16     `json:"endlat"`
	EndLog        int16     `json:"endlog"`
	RiderName     string    `json:"ridername"`
	DriverName    string    `json:"drivername"`
	DriverVehicle string    `json:"drivervehicle"`
	Create        time.Time `json:"create"`
}
