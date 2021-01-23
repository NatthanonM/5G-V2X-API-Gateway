package models

import "time"

// BaseResponse ...
type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// MapResponse ...
type MapResponse struct {
	BaseResponse
	Data *MapResponseData `json:"data"`
}

// MapResponseData ...
type MapResponseData struct {
	Accidents []Accident `json:"accident"`
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
