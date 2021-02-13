package controllers

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/models"
	"5g-v2x-api-gateway-service/internal/services"
	"log"
	"net/http"
	"strconv"
	"time"

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

// WebDrowsinessHeatmap ...
func (r *DrowsinessController) WebDrowsinessHeatmap(c *gin.Context) {
	hour := c.Param("hour")
	hourInt, err := strconv.Atoi(hour)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, models.DrowsinessMapResponse{
			BaseResponse: models.BaseResponse{
				Success: false,
				Message: "Invalid parameter.",
			},
			Data: nil,
		})
		return
	}

	t := time.Now()
	thTimeZone, _ := time.LoadLocation("Asia/Bangkok")
	from := time.Date(t.Year(), t.Month(), t.Day(), hourInt, 0, 0, 0, thTimeZone).UTC()
	to := time.Date(t.Year(), t.Month(), t.Day(), hourInt, 59, 59, 999, thTimeZone).UTC()
	res, err := r.Services.ServiceGateway.DrowsinessService.GetDailyDrowsinessHeatmap(&from, &to)

	if err != nil {
		log.Println(err)
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

// WebAuthDrowsinessMap ...
func (r *DrowsinessController) WebAuthDrowsinessMap(c *gin.Context) {
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

	t := time.Now()
	thTimeZone, _ := time.LoadLocation("Asia/Bangkok")
	from := time.Date(t.Year(), t.Month(), t.Day(), hourInt, 0, 0, 0, thTimeZone).UTC()
	to := time.Date(t.Year(), t.Month(), t.Day(), hourInt, 59, 59, 999, thTimeZone).UTC()
	res, err := r.Services.ServiceGateway.DrowsinessService.GetDailyAuthDrowsinessHeatmap(&from, &to)

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

// WebDrowsinessStatTimebar ...
func (r *DrowsinessController) WebDrowsinessStatTimebar(c *gin.Context) {
	res, err := r.Services.ServiceGateway.DrowsinessService.GetNumberOfDrowsinessTimeBar()

	if err != nil {
		c.JSON(http.StatusBadRequest, models.StatBarResponse{
			BaseResponse: models.BaseResponse{
				Success: false,
				Message: "Invalid parameter.",
			},
			Data: nil,
		})
		return
	}
	if len(res) == 0 {
		c.JSON(http.StatusOK, models.StatBarResponse{
			BaseResponse: models.BaseResponse{
				Success: true,
				Message: "No drowsiness data.",
			},
			Data: res,
		})
		return
	}
	c.JSON(http.StatusOK, models.StatBarResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: "Get drowsiness data successful.",
		},
		Data: res,
	})
}

// WebDrowsinessStatAgebar ...
func (r *DrowsinessController) WebDrowsinessStatAgebar(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, models.BaseResponse{
		Success: false,
		Message: "Not implemented.",
	})
	return
}

// WebDrowsinessStatCalendar ...
func (r *DrowsinessController) WebDrowsinessStatCalendar(c *gin.Context) {
	res, err := r.Services.ServiceGateway.DrowsinessService.GetDrowsinessStatCalendar()

	if err != nil {
		c.JSON(http.StatusBadRequest, models.StatCalResponse{
			BaseResponse: models.BaseResponse{
				Success: false,
				Message: "Invalid parameter.",
			},
			Data: nil,
		})
		return
	}
	if len(res) == 0 {
		c.JSON(http.StatusOK, models.StatCalResponse{
			BaseResponse: models.BaseResponse{
				Success: true,
				Message: "No Drowsiness data.",
			},
			Data: res,
		})
		return
	}
	c.JSON(http.StatusOK, models.StatCalResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: "Get Drowsiness data successful.",
		},
		Data: res,
	})
}

// WebDrowsinessStatGenderpie ...
func (r *DrowsinessController) WebDrowsinessStatGenderpie(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, models.BaseResponse{
		Success: false,
		Message: "Not implemented.",
	})
	return
}
