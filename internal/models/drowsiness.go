package models

import "time"

type Drowsiness struct {
	CarID        string    `json:"car_id"`
	Username     string    `json:"username"`
	Time         time.Time `json:"time"`
	ResponseTime float64   `json:"response_time"`
	WorkingHour  float64   `json:"working_hour"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	Road         string    `json:"road"`
}

//// REST RESPONSE ////

// DrowsinessMapResponse ...
type DrowsinessMapResponse struct {
	BaseResponse
	Data []*PublicDrowsinessData `json:"data"`
}

//// REST RESPONSE ENTITY ////

// PublicDrowsinessData ...
type PublicDrowsinessData struct {
	Detail     AccidentDetail `json:"detail"`
	Coordinate Coordinate     `json:"coordinate"`
}
