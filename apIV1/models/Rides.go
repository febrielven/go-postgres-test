package models

import (
	"time"
)

// Rides schema of rides table ...
type Rides struct {
	ID            int64     `sql:"AUTO_INCREMENT" json:"rideid"`
	StartLat      int16     `json:"startlat"`
	StartLong     int16     `json:"startlong"`
	EndLat        int16     `json:"endlat"`
	EndLong       int16     `json:"endlong"`
	RiderName     string    `json:"ridername"`
	DriverName    string    `json:"drivername"`
	DriverVehicle string    `json:"drivervehicle"`
	Created       time.Time `json:"create"`
}
