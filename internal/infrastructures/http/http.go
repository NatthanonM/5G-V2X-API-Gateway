package http

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/controllers"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
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

	if g.config.Mode != "Development" {
		api.Use(cors.New(cors.Config{
			AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTION"},
			AllowHeaders: []string{"withCredentials", "Access-Control-Allow-Headers", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
			// AllowAllOrigins:  true,
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
			AllowOrigins:     []string{g.config.WebsiteOrigin},
		}))
	}

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
	auth := web.Group("/auth")
	auth.OPTIONS("/login", g.preflight)
	auth.POST("/login", g.Controller.AdminController.WebAuthLogin)
	auth.POST("/logout", g.Controller.AdminController.WebAuthLogout)
	auth.POST("/register", g.Controller.AdminController.WebAuthRegister)
	// auth.GET("/driver", g.Controller.AccidentController.WebAuthGetDriver)
	auth.POST("/driver", g.Controller.DriverController.WebAuthCreateDriver)
	// auth.PATCH("/driver/:id", g.Controller.AccidentController.WebAuthUpdateDriver)
	// auth.DELETE("/driver/:id", g.Controller.AccidentController.WebAuthDeleteDriver)
	auth.GET("/car", g.Controller.CarController.WebAuthGetCar)
	auth.POST("/car", g.Controller.CarController.WebAuthCreateCar)
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

func (g *GinServer) preflight(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", g.config.WebsiteOrigin)
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, struct{}{})
}
