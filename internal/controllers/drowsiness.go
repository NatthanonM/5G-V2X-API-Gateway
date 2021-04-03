package controllers

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/models"
	"5g-v2x-api-gateway-service/internal/services"
	"5g-v2x-api-gateway-service/internal/utils"
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

// WebDrowsinessMap ...
func (r *DrowsinessController) WebDrowsinessMap(c *gin.Context) {
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

	res, err := r.Services.ServiceGateway.DrowsinessService.GetDailyDrowsinessHeatmap(&starttm, &endtm)

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
	res, err := r.Services.ServiceGateway.DrowsinessService.GetDailyAuthDrowsinessHeatmap(&starttm, &endtm)

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

	res, err := r.Services.ServiceGateway.DrowsinessService.GetNumberOfDrowsinessTimeBar(&starttm, &endtm)

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

// WebDrowsinessStatCalendar ...
func (r *DrowsinessController) WebDrowsinessStatCalendar(c *gin.Context) {
	year := c.Query("year")
	i, err := strconv.Atoi(year)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: "Invalid year",
		})
		return
	}

	res, err := r.Services.ServiceGateway.DrowsinessService.GetDrowsinessStatCalendar(int64(i))

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

// WebDrowsinessCount
func (r *DrowsinessController) WebDrowsinessCount(c *gin.Context) {
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

	res, err := r.Services.ServiceGateway.DrowsinessService.GetDrowsinessData(&starttm, &endtm, nil, nil)
	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
		})
		return
	}
	c.JSON(http.StatusOK, models.DrowsinessCountResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: "Get drowsiness count successful.",
		},
		Data: int64(len(res)),
	})
}

// WebAuthDriverDrowsinessStatTimebar ...
func (r *DrowsinessController) WebAuthDriverDrowsinessStatTimebar(c *gin.Context) {
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

	res, err := r.Services.ServiceGateway.DrowsinessService.GetDrowsinessStatTimebar(nil, nil, &driver.Username)
	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
		})
		return
	}
	c.JSON(http.StatusOK, models.WebAuthDriverDrowsinessStatTimebarResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: "Get drowsiness timebar successful.",
		},
		Data: res,
	})
}
