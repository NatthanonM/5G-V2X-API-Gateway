package controllers

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/models"
	"5g-v2x-api-gateway-service/internal/services"
	"5g-v2x-api-gateway-service/internal/utils"
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

	publicAccidentData := []*models.AccidentData{}
	for _, accident := range res {
		publicAccidentData = append(publicAccidentData, &models.AccidentData{
			Detail: models.AccidentDetail{
				Road: accident.Road,
				Time: accident.Time,
			},
			Coordinate: models.Coordinate{
				Lat: accident.Latitude,
				Lng: accident.Longitude,
			},
		})
	}

	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
		})
		return
	}
	if len(res) == 0 {
		c.JSON(http.StatusOK, models.CarAccidentResponse{
			BaseResponse: models.BaseResponse{
				Success: true,
				Message: "No accident data.",
			},
			Data: publicAccidentData,
		})
		return
	}
	c.JSON(http.StatusOK, models.CarAccidentResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: "Get accident data successful.",
		},
		Data: publicAccidentData,
	})
}

// WebAccidentMap ...
func (r *AccidentController) WebAccidentMap(c *gin.Context) {
	start := c.Query("start")
	stop := c.Query("stop")

	i, err := strconv.ParseInt(start, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: "Start date is invalid",
		})
		return
	}
	starttm := time.Unix(i, 0)

	i, err = strconv.ParseInt(stop, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: "Stop date is invalid",
		})
		return
	}
	stoptm := time.Unix(i, 0)

	res, err := r.Services.ServiceGateway.AccidentService.GetDailyAccidentMap(&starttm, &stoptm)

	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
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

// WebAccidentStatCalendar ...
func (r *AccidentController) WebAccidentStatCalendar(c *gin.Context) {
	year := c.Param("year")
	i, err := strconv.Atoi(year)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: "Invalid year",
		})
		return
	}

	res, err := r.Services.ServiceGateway.AccidentService.GetAccidentStatCalendar(int64(i))

	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
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
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
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
	start := c.Query("start")
	stop := c.Query("stop")

	i, err := strconv.ParseInt(start, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: "Start date is invalid",
		})
		return
	}
	starttm := time.Unix(i, 0)

	i, err = strconv.ParseInt(stop, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: "Stop date is invalid",
		})
		return
	}
	stoptm := time.Unix(i, 0)

	res, err := r.Services.ServiceGateway.AccidentService.GetNumberOfAccidentTimeBar(&starttm, &stoptm)

	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
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
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
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

// WebAuthDriverAccidentStatTimebar ...
func (r *AccidentController) WebAuthDriverAccidentStatTimebar(c *gin.Context) {
	driverID := c.Param("id")

	driver, err := r.Services.ServiceGateway.DriverService.GetDriver(driverID)
	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
		})
		return
	}

	res, err := r.Services.ServiceGateway.AccidentService.GetAccidentStatTimebar(nil, nil, &driver.Username)
	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
		})
		return
	}
	c.JSON(http.StatusOK, models.WebAuthDriverAccidentStatTimebarResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: "Get accident timebar successful.",
		},
		Data: res,
	})
}

func (r *AccidentController) WebAccidentStatRoadToptenYear(c *gin.Context) {
	year := c.Param("year")
	i, err := strconv.Atoi(year)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: "Invalid year",
		})
		return
	}

	fromTime := time.Date(i, time.Month(1), 1, 0, 0, 0, 1, time.UTC)
	toTime := time.Date(i, time.Month(12), 31, 23, 59, 59, 999999999, time.UTC)

	var n int64
	n = 10
	res, err := r.Services.ServiceGateway.AccidentService.GetStatRoadToptenYear(&fromTime, &toTime, &n)
	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
		})
		return
	}

	c.JSON(http.StatusOK, models.WebAccidentStatRoadToptenYearResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: "Get accident timebar successful.",
		},
		Data: res,
	})
}
