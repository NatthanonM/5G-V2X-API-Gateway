package models

import "time"

type Accident struct {
	CarID     string    `json:"car_id"`
	Username  string    `json:"username"`
	Time      time.Time `json:"time"`
	Latitude  float64   `json:"lat"`
	Longitude float64   `json:"lng"`
	Road      string    `json:"road"`
}

//// REST RESPONSE ////

// AccidentMapResponse ...
type AccidentMapResponse struct {
	BaseResponse
	Data []*AccidentData `json:"data"`
}

// StatCalResponse ...
type StatCalResponse struct {
	BaseResponse
	Data []*StatCal `json:"data"`
}

// StatBarResponse ...
type StatBarResponse struct {
	BaseResponse
	Data []int32 `json:"data"`
}

// StatPieResponse ...
type StatPieResponse struct {
	BaseResponse
	Data *StatPie `json:"data"`
}

// CarAccidentResponse ...
type CarAccidentResponse struct {
	BaseResponse
	Data []*Accident `json:"data"`
}

//// REST RESPONSE ENTITY ////

// AccidentData ...
type AccidentData struct {
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
	Time   time.Time `json:"time"`
	Road   string    `json:"road"`
	Driver *Driver   `json:"driver"`
	Car    *Car      `json:"car"`
}

// Coordinate ...
type Coordinate struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
