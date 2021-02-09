package middleware

import (
	"5g-v2x-api-gateway-service/internal/models"
	"5g-v2x-api-gateway-service/internal/services"
	"5g-v2x-api-gateway-service/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware ...
type AuthMiddleware struct {
	ServiceGateway services.ServiceGateway
}

// NewAuthMiddleware ...
func NewAuthMiddleware(serviceGateway services.ServiceGateway) *AuthMiddleware {
	return &AuthMiddleware{
		ServiceGateway: serviceGateway,
	}
}

// AuthAdminMiddleware ...
func (a *AuthMiddleware) AuthAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("accessToken")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.BaseResponse{
				Success: false,
				Message: "Authorization invalid"},
			)
			return
		}

		username, err := a.ServiceGateway.AdminService.VerifyAccessToken(accessToken)

		if err != nil {
			customError := utils.NewCustomError(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.BaseResponse{
				Success: false,
				Message: customError.Message},
			)
			return
		}

		c.Set(utils.UsernameCtxKey, *username)
	}
}
