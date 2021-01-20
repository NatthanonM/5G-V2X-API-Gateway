package http

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/controllers"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// GinServer ...
type GinServer struct {
	route      *gin.Engine
	config     *config.Config
	Controller controllers.ControllerGateway
}

// ResolverGateway ...
type ResolverGateway struct {
	dig.In
}

// NewGinServer ...
func NewGinServer(
	cg controllers.ControllerGateway,
	config *config.Config,
) *GinServer {
	h := &GinServer{
		Controller: cg,
		config:     config,
	}
	h.configure()
	return h
}

func (g *GinServer) configure() {
	g.route = gin.Default()

	g.route.GET("/ping", g.Controller.AccidentController.Ping)

}

// Start ...
func (g *GinServer) Start() error {
	return g.route.Run(":" + g.config.Port)
}
