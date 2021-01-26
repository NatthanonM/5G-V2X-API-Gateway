package services

import (
	"go.uber.org/dig"
)

// Service ...
type Service struct {
	ServiceGateway ServiceGateway
}

// ServiceGateway ...
type ServiceGateway struct {
	dig.In
	*AccidentService
	*DrowsinessService
}

// NewService ...
func NewService(sg ServiceGateway) *Service {
	return &Service{
		ServiceGateway: sg,
	}
}
