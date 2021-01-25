package controllers

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/models"
	"5g-v2x-api-gateway-service/internal/services"
	"net/http"

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

func (r *DrowsinessController) WebDrowsinessHeatmap(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, models.MapResponse{
		BaseResponse: models.BaseResponse{
			Success: false,
			Message: "Not implemented.",
		},
		Data: nil,
	})
	return
}

func (r *DrowsinessController) WebDrowsinessStatTimebar(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, models.MapResponse{
		BaseResponse: models.BaseResponse{
			Success: false,
			Message: "Not implemented.",
		},
		Data: nil,
	})
	return
}

func (r *DrowsinessController) WebDrowsinessStatAgebar(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, models.MapResponse{
		BaseResponse: models.BaseResponse{
			Success: false,
			Message: "Not implemented.",
		},
		Data: nil,
	})
	return
}

func (r *DrowsinessController) WebDrowsinessStatCalendar(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, models.MapResponse{
		BaseResponse: models.BaseResponse{
			Success: false,
			Message: "Not implemented.",
		},
		Data: nil,
	})
	return
}

func (r *DrowsinessController) WebDrowsinessStatGenderpie(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, models.MapResponse{
		BaseResponse: models.BaseResponse{
			Success: false,
			Message: "Not implemented.",
		},
		Data: nil,
	})
	return
}
