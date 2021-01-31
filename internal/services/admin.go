package services

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/repositories"
	proto "5g-v2x-api-gateway-service/pkg/api"
)

// AdminService ...
type AdminService struct {
	AdminRepository *repositories.AdminRepository
	config          *config.Config
}

// NewAdminService ...
func NewAdminService(repo *repositories.AdminRepository, cf *config.Config) *AdminService {
	return &AdminService{
		AdminRepository: repo,
		config:          cf,
	}
}

func (as *AdminService) Register(username, password string) error {
	request := proto.RegisterAdminRequest{
		Username: username,
		Password: password,
	}
	if err := as.AdminRepository.Register(&request); err != nil {
		return err
	}
	return nil
}
