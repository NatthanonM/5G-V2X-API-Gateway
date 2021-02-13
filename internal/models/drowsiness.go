package models

import "time"

type Drowsiness struct {
	CarID        string    `json:"car_id"`
	Username     string    `json:"username"`
	Time         time.Time `json:"time"`
	ResponseTime float64   `json:"response_time"`
	WorkingHour  float64   `json:"working_hour"`
	Latitude     float64   `json:"lat"`
	Longitude    float64   `json:"lng"`
	Road         string    `json:"road"`
}

//// REST RESPONSE ////

// DrowsinessMapResponse ...
type DrowsinessMapResponse struct {
	BaseResponse
	Data []*DrowsinessData `json:"data"`
}

type WebAuthDriverDrowsinessStatTimebarResponse struct {
	BaseResponse
	Data []int64 `json:"data"`
}

//// REST RESPONSE ENTITY ////

// DrowsinessData ...
type DrowsinessData struct {
	Detail     DrowsinessDetail `json:"detail"`
	Coordinate Coordinate       `json:"coordinate"`
}

// DrowsinessDetail ...
type DrowsinessDetail struct {
	Time   time.Time `json:"time"`
	Road   string    `json:"road"`
	Driver *Driver   `json:"driver"`
	Car    *Car      `json:"car"`
}
