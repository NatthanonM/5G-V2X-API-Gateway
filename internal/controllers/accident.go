package controllers

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/models"
	"5g-v2x-api-gateway-service/internal/services"
	"net/http"
	"strconv"

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

func (r *AccidentController) WebAccidentMap(c *gin.Context) {
	hour := c.Param("hour")
	hourInt, err := strconv.Atoi(hour)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.MapResponse{
			BaseResponse: models.BaseResponse{
				Success: false,
				Message: "Invalid parameter.",
			},
			Data: nil,
		})
		return
	}

	res, err := r.Services.ServiceGateway.AccidentService.GetDailyAccidentMap(int32(hourInt))

	if err != nil {
		c.JSON(http.StatusBadRequest, models.MapResponse{
			BaseResponse: models.BaseResponse{
				Success: false,
				Message: "Invalid parameter.",
			},
			Data: nil,
		})
		return
	}
	if len(res.Accidents) == 0 {
		c.JSON(http.StatusOK, models.MapResponse{
			BaseResponse: models.BaseResponse{
				Success: true,
				Message: "No accident data.",
			},
			Data: res,
		})
		return
	}
	c.JSON(http.StatusOK, models.MapResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: "Get accident data successful.",
		},
		Data: res,
	})
}

func (r *AccidentController) WebAccidentHeatmap(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, models.MapResponse{
		BaseResponse: models.BaseResponse{
			Success: false,
			Message: "Not implemented.",
		},
		Data: nil,
	})
	return
}

func (r *AccidentController) WebAccidentStatCalendar(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, models.MapResponse{
		BaseResponse: models.BaseResponse{
			Success: false,
			Message: "Not implemented.",
		},
		Data: nil,
	})
	return
}

func (r *AccidentController) WebAccidentStatRoadpie(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, models.MapResponse{
		BaseResponse: models.BaseResponse{
			Success: false,
			Message: "Not implemented.",
		},
		Data: nil,
	})
	return
}

func (r *AccidentController) WebAccidentStatTimebar(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, models.MapResponse{
		BaseResponse: models.BaseResponse{
			Success: false,
			Message: "Not implemented.",
		},
		Data: nil,
	})
	return
}

func (r *AccidentController) WebAccidentStatAgebar(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, models.MapResponse{
		BaseResponse: models.BaseResponse{
			Success: false,
			Message: "Not implemented.",
		},
		Data: nil,
	})
	return
}

func (r *AccidentController) WebAccidentStatGenderbar(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, models.MapResponse{
		BaseResponse: models.BaseResponse{
			Success: false,
			Message: "Not implemented.",
		},
		Data: nil,
	})
	return
}
