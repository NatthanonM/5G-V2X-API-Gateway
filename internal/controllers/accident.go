package controllers

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/models"
	"5g-v2x-api-gateway-service/internal/services"
	"5g-v2x-api-gateway-service/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AccidentController ...
type AccidentController struct {
	Services *services.Service
	config   *config.Config
}

// NewAccidentController ...
func NewAccidentController(srv *services.Service, cf *config.Config) *AccidentController {
	return &AccidentController{
		Services: srv,
		config:   cf,
	}
}

// func (r *AccidentController) Ping(c *gin.Context) {
// 	c.JSON(http.StatusOK, models.PingResponse{
// 		Success: true,
// 		Message: "Pong!",
// 	})
// }

func (r *AccidentController) Map(c *gin.Context) {
	res, err := r.Services.ServiceGateway.AccidentService.GetDailyAccidentMap()
	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.MapResponse{
			BaseResponse: models.BaseResponse{
				Success: false,
				Message: customError.Message,
			},
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, models.MapResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: "Yay!",
		},
		Data: res,
	})
}
