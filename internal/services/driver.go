package services

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/models"
	"5g-v2x-api-gateway-service/internal/repositories"
	"5g-v2x-api-gateway-service/internal/utils"
	proto "5g-v2x-api-gateway-service/pkg/api"
	"time"
)

// DriverService ...
type DriverService struct {
	DriverRepository *repositories.DriverRepository
	config           *config.Config
}

// NewDriverService ...
func NewDriverService(repo *repositories.DriverRepository, cf *config.Config) *DriverService {
	return &DriverService{
		DriverRepository: repo,
		config:           cf,
	}
}

func (ds *DriverService) AddNewDriver(
	firstname, lastname, username, password string,
	dateOfBirth time.Time,
	gender int64) (*string, error) {
	request := proto.AddNewDriverRequest{
		Firstname:   firstname,
		Lastname:    lastname,
		Username:    username,
		Password:    password,
		DateOfBirth: utils.WrapperTime(&dateOfBirth),
		Gender:      proto.AddNewDriverRequest_GenderType(gender),
	}
	res, err := ds.DriverRepository.AddNewDriver(&request)
	if err != nil {
		return nil, err
	}
	return &res.DriverId, nil
}

func (ds *DriverService) GetAllDriver() ([]*models.Driver, error) {
	drivers, err := ds.DriverRepository.GetAllDriver()
	if err != nil {
		return nil, err
	}
	driverList := []*models.Driver{}
	for _, driver := range drivers.Drivers {
		driverList = append(driverList, &models.Driver{DriverID: driver.DriverID,
			Firstname:   driver.Firstname,
			Lastname:    driver.Lastname,
			DateOfBirth: driver.DateOfBirth.AsTime(),
			Gender:      driver.Gender})
	}
	return driverList, nil
}
