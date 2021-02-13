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

	if temp.Username == "" || temp.Password == "" {
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
		temp.Firstname, temp.Lastname, temp.Username, temp.Password, temp.DateOfBirth, genderInt)

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

	privateAccidentData := []*models.Accident{}

	for _, accident := range accidents {
		privateAccidentData = append(privateAccidentData, &models.Accident{
			Time:      accident.Time,
			CarID:     accident.CarID,
			Username:  accident.Username,
			Latitude:  accident.Latitude,
			Longitude: accident.Longitude,
			Road:      accident.Road,
		})
	}

	// Success
	c.JSON(http.StatusOK, models.WebAuthDriverAccidentResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: fmt.Sprintf(`Get accidents of %s successful.`, driver.Username),
		},
		Data: accidents,
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

	drowsinesses, err := dc.Services.ServiceGateway.GetDrowsinessData(nil, &driver.Username)
	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
		})
		return
	}

	privateDrowsinessesData := []*models.Drowsiness{}

	for _, drowsiness := range drowsinesses {
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

	// Success
	c.JSON(http.StatusOK, models.WebAuthDriverDrowsinessResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: fmt.Sprintf(`Get drowsiness of %s successful.`, driver.Username),
		},
		Data: privateDrowsinessesData,
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
