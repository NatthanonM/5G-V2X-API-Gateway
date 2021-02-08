package controllers

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/models"
	"5g-v2x-api-gateway-service/internal/services"
	"5g-v2x-api-gateway-service/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CarController ...
type CarController struct {
	Services *services.Service
	config   *config.Config
}

// NewCarController ...
func NewCarController(srv *services.Service, cf *config.Config) *CarController {
	return &CarController{
		Services: srv,
		config:   cf,
	}
}

func (cc *CarController) WebAuthCreateCar(c *gin.Context) {
	var temp models.NewCarBody
	c.BindJSON(&temp)

	if temp.CarDetail == "" || temp.VehicleRegistrationNumber == "" {
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: "Invalid parameter.",
		})
		return
	}

	carID, err := cc.Services.ServiceGateway.CarService.RegisterNewCar(temp.CarDetail, temp.VehicleRegistrationNumber)

	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
		})
		return
	}

	// Success
	c.JSON(http.StatusOK, models.WebAuthCreateCar{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: "Register car successful.",
		},
		Data: &models.Car{
			CarID: *carID,
		},
	})
}

func (cc *CarController) WebAuthGetCars(c *gin.Context) {
	carList, err := cc.Services.ServiceGateway.CarService.GetCarList()

	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
		})
		return
	}

	// Success
	c.JSON(http.StatusOK, models.WebAuthGetCars{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: "Register car successful.",
		},
		Data: carList,
	})
}

func (cc *CarController) WebAuthGetCar(c *gin.Context) {
	carID := c.Param("id")
	car, err := cc.Services.ServiceGateway.CarService.GetCar(carID)

	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
		})
		return
	}

	accident, err := cc.Services.ServiceGateway.AccidentService.GetAccidentCar(nil, nil, &carID)

	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
		})
		return
	}

	drowsiness, err := cc.Services.ServiceGateway.DrowsinessService.GetDrowsinessData(&carID)

	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
		})
		return
	}

	// Success
	c.JSON(http.StatusOK, models.WebAuthGetCar{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: "Register car successful.",
		},
		Data: &models.WebAuthGetCarResponseData{
			Car:        car,
			Accident:   accident,
			Drowsiness: drowsiness,
		},
	})
}
