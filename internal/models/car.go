package models

import "time"

type Car struct {
	CarID                     string    `json:"car_id"`
	VehicleRegistrationNumber string    `json:"vehicle_registration_number"`
	CarDetail                 string    `json:"car_detail"`
	RegisteredAt              time.Time `json:"registered_at"`
	CreatedAt                 time.Time `json:"created_at"`
}

//// REST BODY ////

// NewCarBody ...
type NewCarBody struct {
	VehicleRegistrationNumber string `json:"vehicle_registration_number"`
	CarDetail                 string `json:"car_detail"`
}

//// REST RESPONSE ////

// WebAuthCreateCar ...
type WebAuthCreateCar struct {
	BaseResponse
	Data *Car `json:"data"`
}

// WebAuthGetCars ...
type WebAuthGetCars struct {
	BaseResponse
	Data []*Car `json:"data"`
}

// WebAuthGetCar ...
type WebAuthGetCar struct {
	BaseResponse
	Data *WebAuthGetCarResponseData `json:"data"`
}

//// REST RESPONSE ENTITY ////

// WebAuthGetCarResponseData ...
type WebAuthGetCarResponseData struct {
	Car        *Car          `json:"car"`
	Accident   []*Accident   `json:"accident"`
	Drowsiness []*Drowsiness `json:"drowsiness"`
}
