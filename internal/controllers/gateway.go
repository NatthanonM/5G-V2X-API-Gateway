package controllers

import (
	"go.uber.org/dig"
)

// Controller ...
type Controller struct {
	ControllerGateway ControllerGateway
}

// ControllerGateway ...
type ControllerGateway struct {
	dig.In
	*AccidentController
	*DrowsinessController
	*AdminController
	*DriverController
	*CarController
}

// NewController ...
func NewController(cg ControllerGateway) *Controller {
	return &Controller{
		ControllerGateway: cg,
	}
}
