package models

import "time"

// BaseResponse ...
type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// AccidentMapResponse ...
type AccidentMapResponse struct {
	BaseResponse
	Data []*Accident `json:"data"`
}
type StatCalResponse struct {
	BaseResponse
	Data []*StatCal `json:"data"`
}
type StatBarResponse struct {
	BaseResponse
	Data []int32 `json:"data"`
}
type StatPieResponse struct {
	BaseResponse
	Data *StatPie `json:"data"`
}

// Accident ...
type Accident struct {
	Detail     AccidentDetail `json:"detail"`
	Coordinate Coordinate     `json:"coordinate"`
}

// StatCal
type StatCal struct {
	Name string  `json:"name"`
	Data []int32 `json:"data"`
}

//statpie
type StatPie struct {
	Series []int32  `json:"series"`
	Labels []string `json:"labels"`
}

// AccidentDetail ...
type AccidentDetail struct {
	Time time.Time `json:"time"`
}

// Coordinate ...
type Coordinate struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// DrowsinessMapResponse ...
type DrowsinessMapResponse struct {
	BaseResponse
	Data []*Drowsiness `json:"data"`
}

// Drowsiness ...
type Drowsiness struct {
	Detail     AccidentDetail `json:"detail"`
	Coordinate Coordinate     `json:"coordinate"`
}

// AdminRegisterBody ...
type AdminRegisterBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AdminLoginBody ...
type AdminLoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// WebAuthCreateDriverResponse ...
type WebAuthCreateDriverResponse struct {
	BaseResponse
	Data *Driver `json:"data"`
}

// NewDriverBody ...
type NewDriverBody struct {
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Gender      int64     `json:"gender"`
}

// NewCarBody ...
type NewCarBody struct {
	VehicleRegistrationNumber string `json:"vehicle_registration_number"`
	CarDetail                 string `json:"car_detail"`
}

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

// WebAuthGetCarResponseData ...
type WebAuthGetCarResponseData struct {
	Car        *Car          `json:"car"`
	Accident   []*Accident   `json:"accident"`
	Drowsiness []*Drowsiness `json:"drowsiness"`
}

// CarAccidentResponse ...
type CarAccidentResponse struct {
	BaseResponse
	Data []*Accident `json:"data"`
}

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
