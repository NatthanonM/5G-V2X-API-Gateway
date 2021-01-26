package container

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/controllers"
	"5g-v2x-api-gateway-service/internal/infrastructures/grpc"
	"5g-v2x-api-gateway-service/internal/infrastructures/http"
	"5g-v2x-api-gateway-service/internal/repositories"
	"5g-v2x-api-gateway-service/internal/services"

	"go.uber.org/dig"
)

// Container ...
type Container struct {
	container *dig.Container
	Error     error
}

// NewContainer ...
func NewContainer() *Container {
	c := &Container{}
	c.configure()
	return c
}

func (cn *Container) configure() {
	cn.container = dig.New()

	// config
	if err := cn.container.Provide(config.NewConfig); err != nil {
		cn.Error = err
	}

	// infrastructure
	if err := cn.container.Provide(http.NewGinServer); err != nil {
		cn.Error = err
	}
	if err := cn.container.Provide(grpc.NewGRPC); err != nil {
		cn.Error = err
	}

	// Service
	if err := cn.container.Provide(services.NewService); err != nil {
		cn.Error = err
	}

	if err := cn.container.Provide(services.NewAccidentService); err != nil {
		cn.Error = err
	}

	if err := cn.container.Provide(services.NewDrowsinessService); err != nil {
		cn.Error = err
	}

	// Repository
	if err := cn.container.Provide(repositories.NewAccidentRepository); err != nil {
		cn.Error = err
	}

	if err := cn.container.Provide(repositories.NewDrowsinessRepository); err != nil {
		cn.Error = err
	}

	// Controller
	if err := cn.container.Provide(controllers.NewController); err != nil {
		cn.Error = err
	}

	if err := cn.container.Provide(controllers.NewAccidentController); err != nil {
		cn.Error = err
	}

	if err := cn.container.Provide(controllers.NewDrowsinessController); err != nil {
		cn.Error = err
	}

}

// Run ...
func (cn *Container) Run() *Container {
	if err := cn.container.Invoke(func(g *http.GinServer) {
		if err := g.Start(); err != nil {
			panic(err)
		}
	}); err != nil {
		panic(err)
	}
	return cn
}
