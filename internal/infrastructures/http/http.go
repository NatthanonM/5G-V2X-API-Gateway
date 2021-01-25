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

	api := g.route.Group("/api")

	// drowsiness := api.Group("/drowsiness")

	//// FOR CAR
	// car := api.Group("/car")
	// car.GET("/accident", g.Controller.AccidentController.CarAccident)
	// car.POST("/login", g.Controller.AccidentController.CarLogin)
	//// FOR WEB
	web := api.Group("/web")
	//// FOR ACCIDENT WEB
	accident := web.Group("/accident")
	accident.GET("/map/:hour", g.Controller.AccidentController.WebAccidentMap)
	accident.GET("/heatmap/:hour", g.Controller.AccidentController.WebAccidentHeatmap)
	accident.GET("/stat/calendar", g.Controller.AccidentController.WebAccidentStatCalendar)
	accident.GET("/stat/roadpie", g.Controller.AccidentController.WebAccidentStatRoadpie)
	accident.GET("/stat/timebar", g.Controller.AccidentController.WebAccidentStatTimebar)
	accident.GET("/stat/agebar", g.Controller.AccidentController.WebAccidentStatAgebar)
	accident.GET("/stat/genderbar", g.Controller.AccidentController.WebAccidentStatGenderbar)
	//// FOR DROWSINESS WEB
	drowsiness := web.Group("/drowsiness")
	drowsiness.GET("/heatmap/:hour", g.Controller.DrowsinessController.WebDrowsinessHeatmap)
	drowsiness.GET("/stat/timebar", g.Controller.DrowsinessController.WebDrowsinessStatTimebar)
	drowsiness.GET("/stat/agebar", g.Controller.DrowsinessController.WebDrowsinessStatAgebar)
	drowsiness.GET("/stat/calendar", g.Controller.DrowsinessController.WebDrowsinessStatCalendar)
	drowsiness.GET("/stat/genderpie", g.Controller.DrowsinessController.WebDrowsinessStatGenderpie)
	//// FOR AUTH WEB
	// auth := web.Group("/auth")
	// auth.POST("/login", g.Controller.AccidentController.WebAuthLogin)
	// auth.POST("/register", g.Controller.AccidentController.WebAuthRegister)
	// auth.GET("/driver", g.Controller.AccidentController.WebAuthGetDriver)
	// auth.POST("/driver", g.Controller.AccidentController.WebAuthCreateDriver)
	// auth.PATCH("/driver/:id", g.Controller.AccidentController.WebAuthUpdateDriver)
	// auth.DELETE("/driver/:id", g.Controller.AccidentController.WebAuthDeleteDriver)
	// auth.GET("/car", g.Controller.AccidentController.WebAuthGetCar)
	// auth.POST("/car", g.Controller.AccidentController.WebAuthCreateCar)
	// auth.PATCH("/car/:id", g.Controller.AccidentController.WebAuthUpdateCar)
	// auth.DELETE("/car/:id", g.Controller.AccidentController.WebAuthDeleteCar)
	// auth.GET("/accident/map/:hour", g.Controller.AccidentController.WebAuthAccidentMap)
	// auth.GET("/drowsiness/map/:hour", g.Controller.AccidentController.WebAuthDrowsinessMap)
	// auth.GET("/driver/:id", g.Controller.AccidentController.WebAuthDriver)
	// auth.GET("/driver/:id/accident", g.Controller.AccidentController.WebAuthDriverAccident)
	// auth.GET("/driver/:id/accident/stat", g.Controller.AccidentController.WebAuthDriverAccidentStat)
	// auth.GET("/driver/:id/drowsiness", g.Controller.AccidentController.WebAuthDriverDrowsiness)
	// auth.GET("/driver/:id/drowsiness/stat", g.Controller.AccidentController.WebAuthDriverDrowsinessStat)

}

// Start ...
func (g *GinServer) Start() error {
	return g.route.Run(":" + g.config.Port)
}
