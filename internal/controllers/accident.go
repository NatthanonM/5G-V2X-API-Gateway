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
	from := time.Now().Add(-(60 * time.Minute)).Unix()
	to := time.Now().Unix()
	starttm := time.Unix(from, 0)
	endtm := time.Unix(to, 0)

	res, err := r.Services.ServiceGateway.AccidentService.GetAccident(&starttm, &endtm, nil, nil)

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
	end := c.Query("end")

	i, err := strconv.ParseInt(start, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: "Start date is invalid",
		})
		return
	}
	starttm := time.Unix(i, 0)

	i, err = strconv.ParseInt(end, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: "End date is invalid",
		})
		return
	}
	endtm := time.Unix(i, 0)

	res, err := r.Services.ServiceGateway.AccidentService.GetDailyAccidentMap(&starttm, &endtm)

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
	year := c.Query("year")
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
	end := c.Query("end")

	i, err := strconv.ParseInt(start, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: "Start date is invalid",
		})
		return
	}
	starttm := time.Unix(i, 0).UTC()

	i, err = strconv.ParseInt(end, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: "End date is invalid",
		})
		return
	}
	endtm := time.Unix(i, 0).UTC()

	res, err := r.Services.ServiceGateway.AccidentService.GetNumberOfAccidentTimeBar(&starttm, &endtm)

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

// WebAccidentCount
func (r *AccidentController) WebAccidentCount(c *gin.Context) {
	start := c.Query("date")
	mode := c.Query("mode")

	i, err := strconv.ParseInt(start, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: "Date is invalid",
		})
		return
	}
	starttm := time.Unix(i, 0).UTC()

	var endtm time.Time

	switch mode {
	case "date":
		endtm = time.Date(starttm.Year(), starttm.Month(), starttm.Day()+1,
			starttm.Hour(), starttm.Minute(), starttm.Second(), starttm.Nanosecond(), time.UTC)
	case "week":
		endtm = time.Date(starttm.Year(), starttm.Month(), starttm.Day()+7,
			starttm.Hour(), starttm.Minute(), starttm.Second(), starttm.Nanosecond(), time.UTC)
	case "month":
		endtm = time.Date(starttm.Year(), starttm.Month()+1, starttm.Day(),
			starttm.Hour(), starttm.Minute(), starttm.Second(), starttm.Nanosecond(), time.UTC)
	case "quarter":
		endtm = time.Date(starttm.Year(), starttm.Month()+3, starttm.Day(),
			starttm.Hour(), starttm.Minute(), starttm.Second(), starttm.Nanosecond(), time.UTC)
	case "year":
		endtm = time.Date(starttm.Year()+1, starttm.Month(), starttm.Day(),
			starttm.Hour(), starttm.Minute(), starttm.Second(), starttm.Nanosecond(), time.UTC)
	default:
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: "Mode is invalid",
		})
		return
	}

	res, err := r.Services.ServiceGateway.AccidentService.GetAccident(&starttm, &endtm, nil, nil)
	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
		})
		return
	}
	c.JSON(http.StatusOK, models.AccidentCountResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: "Get accident count successful.",
		},
		Data: int64(len(res)),
	})
}

// WebAuthAccidentMap ...
func (r *AccidentController) WebAuthAccidentMap(c *gin.Context) {
	start := c.Query("start")
	end := c.Query("end")

	i, err := strconv.ParseInt(start, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: "Start date is invalid",
		})
		return
	}
	starttm := time.Unix(i, 0).UTC()

	i, err = strconv.ParseInt(end, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: "End date is invalid",
		})
		return
	}
	endtm := time.Unix(i, 0).UTC()
	res, err := r.Services.ServiceGateway.AccidentService.GetDailyAuthAccidentMap(&starttm, &endtm)

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
	year := c.Query("year")
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
