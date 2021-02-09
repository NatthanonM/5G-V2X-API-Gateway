package models

import "time"

type Driver struct {
	DriverID    string    `json:"driver_id"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Gender      string    `json:"gender"`
	Username    string    `json:"username"`
}
