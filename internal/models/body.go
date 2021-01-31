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

// Accident ...
type Accident struct {
	Detail     AccidentDetail     `json:"detail"`
	Coordinate AccidentCoordinate `json:"coordinate"`
}

// AccidentDetail ...
type AccidentDetail struct {
	Time time.Time `json:"time"`
}

// AccidentCoordinate ...
type AccidentCoordinate struct {
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
	Detail     AccidentDetail     `json:"detail"`
	Coordinate AccidentCoordinate `json:"coordinate"`
}

// AdminRegisterBody ...
type AdminRegisterBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
