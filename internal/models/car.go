package models

import "time"

type Car struct {
	CarID                     string    `json:"car_id"`
	VehicleRegistrationNumber string    `json:"vehicle_registration_number"`
	CarDetail                 string    `json:"car_detail"`
	RegisteredAt              time.Time `json:"registered_at"`
	CreatedAt                 time.Time `json:"created_at"`
}
