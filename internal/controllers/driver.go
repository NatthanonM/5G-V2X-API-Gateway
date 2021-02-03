package controllers

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/models"
	"5g-v2x-api-gateway-service/internal/services"
	"5g-v2x-api-gateway-service/internal/utils"
	"net/http"

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

func (dc *DriverController) WebAuthCreateDriver(c *gin.Context) {
	var temp models.NewDriverBody
	c.BindJSON(&temp)

	if temp.Username == "" || temp.Password == "" {
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: "Invalid parameter.",
		})
		return
	}

	driverID, err := dc.Services.ServiceGateway.DriverService.AddNewDriver(
		temp.Firstname, temp.Lastname, temp.Username, temp.Password, temp.DateOfBirth, temp.Gender)

	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
		})
		return
	}

	// Success
	c.JSON(http.StatusOK, models.WebAuthCreateDriverResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: "Create driver successful.",
		},
		Data: &models.Driver{
			DriverID: *driverID,
		},
	})
}
