package controllers

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/models"
	"5g-v2x-api-gateway-service/internal/services"
	"net/http"
	"strconv"
	"time"

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

// CarAccident ...
func (r *AccidentController) CarAccident(c *gin.Context) {
	from := time.Now().Add(-(60 * time.Minute))
	to := time.Now()

	res, err := r.Services.ServiceGateway.AccidentService.GetAccident(&from, &to, nil, nil)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.CarAccidentResponse{
			BaseResponse: models.BaseResponse{
				Success: false,
				Message: "Internal error.",
			},
			Data: nil,
		})
		return
	}
	if len(res) == 0 {
		c.JSON(http.StatusOK, models.CarAccidentResponse{
			BaseResponse: models.BaseResponse{
				Success: true,
				Message: "No accident data.",
			},
			Data: res,
		})
		return
	}
	c.JSON(http.StatusOK, models.CarAccidentResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: "Get accident data successful.",
		},
		Data: res,
	})
}

// WebAccidentMap ...
func (r *AccidentController) WebAccidentMap(c *gin.Context) {
	hour := c.Param("hour")
	hourInt, err := strconv.Atoi(hour)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.AccidentMapResponse{
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
	res, err := r.Services.ServiceGateway.AccidentService.GetDailyAccidentMap(&from, &to)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.AccidentMapResponse{
			BaseResponse: models.BaseResponse{
				Success: false,
				Message: "Invalid parameter.",
			},
			Data: nil,
		})
		return
	}
	if len(res) == 0 {
		c.JSON(http.StatusOK, models.AccidentMapResponse{
			BaseResponse: models.BaseResponse{
				Success: true,
				Message: "No accident data.",
			},
			Data: res,
		})
		return
	}
	c.JSON(http.StatusOK, models.AccidentMapResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: "Get accident data successful.",
		},
		Data: res,
	})
}

// WebAccidentHeatmap ...
func (r *AccidentController) WebAccidentHeatmap(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, models.BaseResponse{
		Success: false,
		Message: "Not implemented.",
	})
	return
}

// WebAccidentStatCalendar ...
func (r *AccidentController) WebAccidentStatCalendar(c *gin.Context) {

	res, err := r.Services.ServiceGateway.AccidentService.GetAccidentStatCalendar()

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
				Message: "No accident data.",
			},
			Data: res,
		})
		return
	}
	c.JSON(http.StatusOK, models.StatCalResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: "Get accident data successful.",
		},
		Data: res,
	})
}

// WebAccidentStatRoadpie ...
func (r *AccidentController) WebAccidentStatRoadpie(c *gin.Context) {
	res, err := r.Services.ServiceGateway.AccidentService.GetNumberOfAccidentStreet()

	if err != nil {
		c.JSON(http.StatusBadRequest, models.StatPieResponse{
			BaseResponse: models.BaseResponse{
				Success: false,
				Message: "Invalid parameter.",
			},
			Data: nil,
		})
		return
	}
	if res.Series == nil {
		c.JSON(http.StatusOK, models.StatPieResponse{
			BaseResponse: models.BaseResponse{
				Success: true,
				Message: "No accident data.",
			},
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, models.StatPieResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: "Get accident data successful.",
		},
		Data: res,
	})
}

// WebAccidentStatTimebar ...
func (r *AccidentController) WebAccidentStatTimebar(c *gin.Context) {
	res, err := r.Services.ServiceGateway.AccidentService.GetNumberOfAccidentTimeBar()

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
				Message: "No accident data.",
			},
			Data: res,
		})
		return
	}
	c.JSON(http.StatusOK, models.StatBarResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: "Get accident data successful.",
		},
		Data: res,
	})
}

// WebAccidentStatAgebar ...
func (r *AccidentController) WebAccidentStatAgebar(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, models.BaseResponse{
		Success: false,
		Message: "Not implemented.",
	})
	return
}

// WebAccidentStatGenderbar ...
func (r *AccidentController) WebAccidentStatGenderbar(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, models.BaseResponse{
		Success: false,
		Message: "Not implemented.",
	})
	return
}

// WebAuthAccidentMap ...
func (r *AccidentController) WebAuthAccidentMap(c *gin.Context) {
	hour := c.Param("hour")
	hourInt, err := strconv.Atoi(hour)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.AccidentMapResponse{
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
	res, err := r.Services.ServiceGateway.AccidentService.GetDailyAuthAccidentMap(&from, &to)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.AccidentMapResponse{
			BaseResponse: models.BaseResponse{
				Success: false,
				Message: "Invalid parameter.",
			},
			Data: nil,
		})
		return
	}
	if len(res) == 0 {
		c.JSON(http.StatusOK, models.AccidentMapResponse{
			BaseResponse: models.BaseResponse{
				Success: true,
				Message: "No accident data.",
			},
			Data: res,
		})
		return
	}
	c.JSON(http.StatusOK, models.AccidentMapResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: "Get accident data successful.",
		},
		Data: res,
	})
}
