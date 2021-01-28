package controllers

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/models"
	"5g-v2x-api-gateway-service/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DrowsinessController ...
type DrowsinessController struct {
	Services *services.Service
	config   *config.Config
}

// NewDrowsinessController ...
func NewDrowsinessController(srv *services.Service, cf *config.Config) *DrowsinessController {
	return &DrowsinessController{
		Services: srv,
		config:   cf,
	}
}

func (r *DrowsinessController) WebDrowsinessHeatmap(c *gin.Context) {
	hour := c.Param("hour")
	hourInt, err := strconv.Atoi(hour)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DrowsinessMapResponse{
			BaseResponse: models.BaseResponse{
				Success: false,
				Message: "Invalid parameter.",
			},
			Data: nil,
		})
		return
	}

	res, err := r.Services.ServiceGateway.DrowsinessService.GetDailyDrowsinessHeatmap(int32(hourInt))

	if err != nil {
		c.JSON(http.StatusBadRequest, models.DrowsinessMapResponse{
			BaseResponse: models.BaseResponse{
				Success: false,
				Message: "Invalid parameter.",
			},
			Data: nil,
		})
		return
	}
	if len(res) == 0 {
		c.JSON(http.StatusOK, models.DrowsinessMapResponse{
			BaseResponse: models.BaseResponse{
				Success: true,
				Message: "No drowsiness data.",
			},
			Data: res,
		})
		return
	}
	c.JSON(http.StatusOK, models.DrowsinessMapResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: "Get drowsiness data successful.",
		},
		Data: res,
	})
}

func (r *DrowsinessController) WebDrowsinessStatTimebar(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, models.BaseResponse{
		Success: false,
		Message: "Not implemented.",
	})
	return
}

func (r *DrowsinessController) WebDrowsinessStatAgebar(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, models.BaseResponse{
		Success: false,
		Message: "Not implemented.",
	})
	return
}

func (r *DrowsinessController) WebDrowsinessStatCalendar(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, models.BaseResponse{
		Success: false,
		Message: "Not implemented.",
	})
	return
}

func (r *DrowsinessController) WebDrowsinessStatGenderpie(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, models.BaseResponse{
		Success: false,
		Message: "Not implemented.",
	})
	return
}