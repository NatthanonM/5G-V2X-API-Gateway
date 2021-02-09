package models

// BaseResponse ...
type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
