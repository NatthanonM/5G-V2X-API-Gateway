package controllers

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/models"
	"5g-v2x-api-gateway-service/internal/services"
	"5g-v2x-api-gateway-service/internal/utils"
	"net/http"
	"time"

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

func (r *AdminController) setAccessTokenCookie(accessToken string, c *gin.Context) {
	lifetimeDuration, _ := time.ParseDuration(r.config.AccessTokenLifetime)
	// convert duration(ns) -> seconds(s)
	lifetimeDurationSeconds := int(lifetimeDuration.Seconds())
	c.SetCookie("accessToken", accessToken, lifetimeDurationSeconds, "/", r.config.WebsiteDomain, false, true)
}

func (ac *AdminController) WebAuthRegister(c *gin.Context) {
	var temp models.AdminRegisterBody
	c.BindJSON(&temp)

	if temp.Username == "" || temp.Password == "" {
		c.JSON(http.StatusBadRequest,
			models.BaseResponse{
				Success: false,
				Message: "Invalid parameter.",
			})
		return
	}

	if err := ac.Services.ServiceGateway.AdminService.Register(temp.Username, temp.Password); err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
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

func (ac *AdminController) WebAuthLogin(c *gin.Context) {
	var temp models.AdminLoginBody
	c.BindJSON(&temp)

	if temp.Username == "" || temp.Password == "" {
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: "Invalid parameter.",
		})
		return
	}

	accessToken, err := ac.Services.ServiceGateway.AdminService.Login(temp.Username, temp.Password)

	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
		})
		return
	}

	ac.setAccessTokenCookie(*accessToken, c)

	// Success
	c.JSON(http.StatusOK, models.BaseResponse{
		Success: true,
		Message: "login successful",
	})
}

func (ac *AdminController) WebAuthLogout(c *gin.Context) {
	c.SetCookie("accessToken", "", -1, "/", ac.config.WebsiteDomain, false, true)

	// Success
	c.JSON(http.StatusOK, models.BaseResponse{
		Success: true,
		Message: "logout successful",
	})
}

func (ac *AdminController) WebAuthProfile(c *gin.Context) {

	ctxData, _ := c.Get(utils.UsernameCtxKey)
	d, ok := ctxData.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, models.BaseResponse{
			Success: false,
			Message: "Authorization invalid",
		},
		)
		return
	}

	admin, err := ac.Services.ServiceGateway.AdminService.GetProfile(d)

	if err != nil {
		customError := utils.NewCustomError(err)
		c.JSON(http.StatusBadRequest, models.BaseResponse{
			Success: false,
			Message: customError.Message,
		})
		return
	}

	// Success
	c.JSON(http.StatusOK, models.WebAuthProfileResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Message: "Get profile successful",
		},
		Data: admin,
	})
}
