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

//// REST BODY ////

// NewDriverBody ...
type NewDriverBody struct {
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Gender      int64     `json:"gender"`
}

//// REST RESPONSE ////

// WebAuthGetDriversResponse ...
type WebAuthGetDriversResponse struct {
	BaseResponse
	Data []*Driver `json:"data"`
}

// WebAuthGetDriverResponse ...
type WebAuthGetDriverResponse struct {
	BaseResponse
	Data *Driver `json:"data"`
}

// WebAuthCreateDriverResponse ...
type WebAuthCreateDriverResponse struct {
	BaseResponse
	Data *Driver `json:"data"`
}

//// REST RESPONSE ENTITY ////

// WebAuthDriverAccident ...
type WebAuthDriverAccident struct {
	BaseResponse
	Data []*Accident `json:"data"`
}
