package http

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/controllers"
	"5g-v2x-api-gateway-service/internal/infrastructures/middleware"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// GinServer ...
type GinServer struct {
	route          *gin.Engine
	config         *config.Config
	AuthMiddleware *middleware.AuthMiddleware
	Controller     controllers.ControllerGateway
}

// ResolverGateway ...
type ResolverGateway struct {
	dig.In
}

// NewGinServer ...
func NewGinServer(
	cg controllers.ControllerGateway,
	auth *middleware.AuthMiddleware,
	config *config.Config,
) *GinServer {
	h := &GinServer{
		Controller:     cg,
		AuthMiddleware: auth,
		config:         config,
	}
	h.configure()
	return h
}

func (g *GinServer) configure() {
	g.route = gin.Default()

	if g.config.Mode != "Development" {
		g.route.Use(cors.New(cors.Config{
			AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTION"},
			AllowHeaders: []string{"withCredentials", "Access-Control-Allow-Headers", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
			// AllowAllOrigins:  true,
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
			AllowOrigins:     []string{g.config.WebsiteOrigin},
		}))
	}
	api := g.route.Group("/api")

	//// FOR CAR
	car := api.Group("/car")
	car.GET("/accident", g.Controller.AccidentController.CarAccident)
	car.POST("/login", g.Controller.DriverController.CarLogin)
	car.OPTIONS("/login", g.preflight)
	//// FOR WEB
	web := api.Group("/web")
	//// FOR ACCIDENT WEB
	accident := web.Group("/accident")
	accident.GET("/map", g.Controller.AccidentController.WebAccidentMap)
	accident.GET("/stat/calendar", g.Controller.AccidentController.WebAccidentStatCalendar)
	accident.GET("/stat/roadpie", g.Controller.AccidentController.WebAccidentStatRoadpie)
	accident.GET("/stat/road/topten", g.Controller.AccidentController.WebAccidentStatRoadToptenYear)
	accident.GET("/stat/timebar", g.Controller.AccidentController.WebAccidentStatTimebar)
	accident.GET("/stat/count", g.Controller.AccidentController.WebAccidentCount)
	//// FOR DROWSINESS WEB
	drowsiness := web.Group("/drowsiness")
	drowsiness.GET("/map", g.Controller.DrowsinessController.WebDrowsinessMap)
	drowsiness.GET("/stat/timebar", g.Controller.DrowsinessController.WebDrowsinessStatTimebar)
	drowsiness.GET("/stat/calendar", g.Controller.DrowsinessController.WebDrowsinessStatCalendar)
	drowsiness.GET("/stat/count", g.Controller.DrowsinessController.WebDrowsinessCount)
	//// FOR AUTH WEB
	auth := web.Group("/auth")
	auth.OPTIONS("/login", g.preflight)
	auth.POST("/login", g.Controller.AdminController.WebAuthLogin)
	auth.POST("/register", g.Controller.AdminController.WebAuthRegister)
	auth.Use(g.AuthMiddleware.AuthAdminMiddleware())
	{
		auth.GET("/profile", g.Controller.AdminController.WebAuthProfile)
		auth.POST("/logout", g.Controller.AdminController.WebAuthLogout)
		auth.OPTIONS("/logout", g.preflight)
		auth.POST("/driver", g.Controller.DriverController.WebAuthCreateDriver)
		auth.OPTIONS("/driver", g.preflight)
		auth.GET("/driver", g.Controller.DriverController.WebAuthGetDrivers)
		auth.GET("/driver/:id", g.Controller.DriverController.WebAuthGetDriver)
		auth.PATCH("/driver/:id", g.Controller.DriverController.WebAuthUpdateDriver)
		auth.DELETE("/driver/:id", g.Controller.DriverController.WebAuthDeleteDriver)
		auth.GET("/car", g.Controller.CarController.WebAuthGetCars)
		auth.GET("/car/:id", g.Controller.CarController.WebAuthGetCar)
		auth.POST("/car", g.Controller.CarController.WebAuthCreateCar)
		auth.OPTIONS("/car", g.preflight)
		// auth.PATCH("/car/:id", g.Controller.AccidentController.WebAuthUpdateCar)
		// auth.DELETE("/car/:id", g.Controller.AccidentController.WebAuthDeleteCar)
		auth.GET("/accident/map", g.Controller.AccidentController.WebAuthAccidentMap)
		auth.GET("/drowsiness/map", g.Controller.DrowsinessController.WebAuthDrowsinessMap)
		auth.GET("/driver/:id/accident", g.Controller.DriverController.WebAuthDriverAccident)
		auth.GET("/driver/:id/accident/stat/timebar", g.Controller.AccidentController.WebAuthDriverAccidentStatTimebar)
		auth.GET("/driver/:id/drowsiness", g.Controller.DriverController.WebAuthDriverDrowsiness)
		auth.GET("/driver/:id/drowsiness/stat/timebar", g.Controller.DrowsinessController.WebAuthDriverDrowsinessStatTimebar)
	}

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
