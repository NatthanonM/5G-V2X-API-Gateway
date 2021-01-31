package controllers

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/models"
	"5g-v2x-api-gateway-service/internal/services"
	"5g-v2x-api-gateway-service/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminController ...
type AdminController struct {
	Services *services.Service
	config   *config.Config
}

// NewAdminController ...
func NewAdminController(srv *services.Service, cf *config.Config) *AdminController {
	return &AdminController{
		Services: srv,
		config:   cf,
	}
}

func (ac *AdminController) WebAuthRegister(c *gin.Context) {
	var temp models.AdminRegisterBody
	c.BindJSON(&temp)

	if temp.Username == "" || temp.Password == "" {
		c.JSON(http.StatusBadRequest, models.AccidentMapResponse{
			BaseResponse: models.BaseResponse{
				Success: false,
				Message: "Invalid parameter.",
			},
			Data: nil,
		})
		return
	}

	if err := ac.Services.ServiceGateway.AdminService.Register(temp.Username, temp.Password); err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.AccidentMapResponse{
			BaseResponse: models.BaseResponse{
				Success: false,
				Message: customError.Message,
			},
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusBadRequest, models.AccidentMapResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: "Register successful.",
		},
		Data: nil,
	})
}
