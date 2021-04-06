package controllers

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/models"
	"5g-v2x-api-gateway-service/internal/services"
	"5g-v2x-api-gateway-service/internal/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// DriverController ...
type DriverController struct {
	Services *services.Service
	config   *config.Config
}

// NewDriverController ...
func NewDriverController(srv *services.Service, cf *config.Config) *DriverController {
	return &DriverController{
		Services: srv,
		config:   cf,
	}
}

// WebAuthCreateDriver ...
func (dc *DriverController) WebAuthCreateDriver(c *gin.Context) {
	var temp models.NewDriverBody
	err := c.BindJSON(&temp)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: "Invalid parameter.",
		})
		return
	}

	if temp.Username == "" || temp.Password == "" || temp.DateOfBirth == nil {
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: "Invalid parameter.",
		})
		return
	}
	genderInt, err := strconv.Atoi(temp.Gender)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: "Invalid parameter: gender.",
		})
		return
	}

	driverID, err := dc.Services.ServiceGateway.DriverService.AddNewDriver(
		temp.Firstname, temp.Lastname, temp.Username, temp.Password, *temp.DateOfBirth, genderInt)

	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
		})
		return
	}

	// Success
	c.JSON(http.StatusCreated, models.WebAuthCreateDriverResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: "Create driver successful.",
		},
		Data: &models.Driver{
			DriverID: *driverID,
		},
	})
}

// WebAuthGetDrivers ...
func (dc *DriverController) WebAuthGetDrivers(c *gin.Context) {
	drivers, err := dc.Services.ServiceGateway.DriverService.GetAllDriver()
	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
		})
		return
	}

	// Success
	c.JSON(http.StatusOK, models.WebAuthGetDriversResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: "Get all driver successful.",
		},
		Data: drivers,
	})
}

// WebAuthGetDriver ...
func (dc *DriverController) WebAuthGetDriver(c *gin.Context) {
	driverID := c.Param("id")
	driver, err := dc.Services.ServiceGateway.DriverService.GetDriver(driverID)
	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
		})
		return
	}

	// Success
	c.JSON(http.StatusOK, models.WebAuthGetDriverResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: "Get driver successful.",
		},
		Data: driver,
	})
}

// WebAuthDriverAccident ...
func (dc *DriverController) WebAuthDriverAccident(c *gin.Context) {
	driverID := c.Param("id")

	driver, err := dc.Services.ServiceGateway.DriverService.GetDriver(driverID)
	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
		})
		return
	}

	accidents, err := dc.Services.ServiceGateway.AccidentService.GetAccident(nil, nil, nil, &driver.Username)
	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
		})
		return
	}

	privateAccidentDatas := []*struct {
		Accident *models.Accident `json:"accident"`
		Car      *models.Car      `json:"car"`
	}{}

	carSet := make(map[string]*models.Car)

	for _, accident := range accidents {
		privateAccidentData := struct {
			Accident *models.Accident `json:"accident"`
			Car      *models.Car      `json:"car"`
		}{}
		privateAccidentData.Accident = &models.Accident{
			Time:      accident.Time,
			CarID:     accident.CarID,
			Username:  accident.Username,
			Latitude:  accident.Latitude,
			Longitude: accident.Longitude,
			Road:      accident.Road,
		}
		_, exists := carSet[accident.CarID]
		if !exists {
			if car, err := dc.Services.ServiceGateway.CarService.GetCar(accident.CarID); err == nil {
				carSet[accident.CarID] = car
			} else {
				carSet[accident.CarID] = &models.Car{
					CarID:                     accident.CarID,
					VehicleRegistrationNumber: "Unknown car",
					CarDetail:                 "Unknown car",
				}
			}
		}
		privateAccidentData.Car = &models.Car{
			CarID:                     accident.CarID,
			VehicleRegistrationNumber: carSet[accident.CarID].VehicleRegistrationNumber,
			CarDetail:                 carSet[accident.CarID].CarDetail,
			RegisteredAt:              carSet[accident.CarID].RegisteredAt,
			MfgAt:                     carSet[accident.CarID].MfgAt,
		}
		privateAccidentDatas = append(privateAccidentDatas, &privateAccidentData)
	}

	// Success
	c.JSON(http.StatusOK, models.WebAuthDriverAccidentResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: fmt.Sprintf(`Get accidents of %s successful.`, driver.Username),
		},
		Data: privateAccidentDatas,
	})
}

// WebAuthDriverDrowsiness ...
func (dc *DriverController) WebAuthDriverDrowsiness(c *gin.Context) {
	driverID := c.Param("id")

	driver, err := dc.Services.ServiceGateway.DriverService.GetDriver(driverID)
	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
		})
		return
	}

	drowsinesses, err := dc.Services.ServiceGateway.GetDrowsinessData(nil, nil, nil, &driver.Username)
	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
		})
		return
	}

	var sum1stDivingHour, sumResponseTime, avg1stDivingHour, avgResponseTime float64
	sum1stDivingHour, sumResponseTime, avg1stDivingHour, avgResponseTime = 0, 0, 0, 0
	drowsyDivingHourEachDay := make(map[time.Time]float64)
	privateDrowsinessesData := []*models.Drowsiness{}

	for _, drowsiness := range drowsinesses {
		date := time.Date(drowsiness.Time.Year(), drowsiness.Time.Month(), drowsiness.Time.Day(), 0, 0, 0, 0, time.UTC)
		drowsyDivingHourEachDay[date] = drowsiness.WorkingHour
		sumResponseTime += drowsiness.ResponseTime
		privateDrowsinessesData = append(privateDrowsinessesData, &models.Drowsiness{
			CarID:        drowsiness.CarID,
			Username:     drowsiness.Username,
			Time:         drowsiness.Time,
			ResponseTime: drowsiness.ResponseTime,
			WorkingHour:  drowsiness.WorkingHour,
			Latitude:     drowsiness.Latitude,
			Longitude:    drowsiness.Longitude,
			Road:         drowsiness.Road,
		})
	}

	for _, v := range drowsyDivingHourEachDay {
		sum1stDivingHour += v
	}

	if len(drowsyDivingHourEachDay) != 0 {
		avg1stDivingHour = sum1stDivingHour / float64(len(drowsyDivingHourEachDay))
	}
	if len(drowsinesses) != 0 {
		avgResponseTime = sumResponseTime / float64(len(drowsinesses))
	}

	// Success
	c.JSON(http.StatusOK, models.WebAuthDriverDrowsinessResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: fmt.Sprintf(`Get drowsiness of %s successful.`, driver.Username),
		},
		Data: models.AuthDriverDrowsiness{
			Drowsiness:        privateDrowsinessesData,
			Avg1stDrivingHour: avg1stDivingHour,
			AvgResponse:       avgResponseTime,
		},
	})
}

func (dc *DriverController) CarLogin(c *gin.Context) {
	var temp models.DriverLoginBody
	c.BindJSON(&temp)

	if temp.Username == "" || temp.Password == "" || temp.CarID == "" {
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: "Invalid parameter.",
		})
		return
	}

	_, err := dc.Services.ServiceGateway.DriverService.Login(temp.Username, temp.Password, temp.CarID)

	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
		})
		return
	}

	// Success
	c.JSON(http.StatusCreated, models.BaseResponse{
		Success: true,
		Message: "login successful",
	})
}

// WebAuthUpdateDriver
func (dc *DriverController) WebAuthUpdateDriver(c *gin.Context) {
	driverID := c.Param("id")

	var temp models.DriverUpdateBody
	c.BindJSON(&temp)

	if temp.Firstname == nil && temp.Lastname == nil && temp.DateOfBirth == nil {
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: "Invalid parameter.",
		})
		return
	}

	err := dc.Services.ServiceGateway.DriverService.Update(driverID, temp.Firstname, temp.Lastname, temp.DateOfBirth)

	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
		})
		return
	}

	// Success
	c.JSON(http.StatusCreated, models.BaseResponse{
		Success: true,
		Message: fmt.Sprintf("update driver `%s` successful", driverID),
	})

}

func (dc *DriverController) WebAuthDeleteDriver(c *gin.Context) {
	driverID := c.Param("id")

	err := dc.Services.ServiceGateway.DriverService.Delete(driverID)

	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
		})
		return
	}

	// Success
	c.JSON(http.StatusCreated, models.BaseResponse{
		Success: true,
		Message: fmt.Sprintf("delete driver `%s` successful", driverID),
	})

}
