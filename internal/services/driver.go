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
	*AccidentService
	*DrowsinessService
	config *config.Config
}

// NewDriverService ...
func NewDriverService(repo *repositories.DriverRepository, accSrv *AccidentService, drowsSrv *DrowsinessService, cf *config.Config) *DriverService {
	return &DriverService{
		DriverRepository:  repo,
		config:            cf,
		AccidentService:   accSrv,
		DrowsinessService: drowsSrv,
	}
}

// AddNewDriver ...
func (ds *DriverService) AddNewDriver(
	firstname, lastname, username, password string,
	dateOfBirth time.Time,
	gender int) (*string, error) {
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

// GetAllDriver ...
func (ds *DriverService) GetAllDriver() ([]*models.Driver, error) {
	drivers, err := ds.DriverRepository.GetAllDriver()
	if err != nil {
		return nil, err
	}
	driverList := []*models.Driver{}
	for _, driver := range drivers.Drivers {
		accidents, _ := ds.AccidentService.GetAccident(nil, nil, nil, &driver.Username)
		drowsinesses, _ := ds.DrowsinessService.GetDrowsinessData(nil, nil, nil, &driver.Username)
		var avgResponseTime *float64
		if len(drowsinesses) > 0 {
			var avg float64
			for _, drowsiness := range drowsinesses {
				avg += drowsiness.ResponseTime
			}
			avg = avg / float64(len(drowsinesses))
			avgResponseTime = &avg
		}
		driverList = append(driverList, &models.Driver{
			DriverID:        driver.DriverId,
			Firstname:       driver.Firstname,
			Lastname:        driver.Lastname,
			DateOfBirth:     driver.DateOfBirth.AsTime(),
			Gender:          driver.Gender,
			Username:        driver.Username,
			AccidentCount:   int64(len(accidents)),
			DrowsinessCount: int64(len(drowsinesses)),
			AvgResponseTime: avgResponseTime,
		})
	}
	return driverList, nil
}

// GetDriver ...
func (ds *DriverService) GetDriver(driverID string) (*models.Driver, error) {
	driver, err := ds.DriverRepository.GetDriver(&proto.GetDriverRequest{
		DriverId: driverID,
	})
	if err != nil {
		return nil, err
	}
	return &models.Driver{
		DriverID:    driver.DriverId,
		Firstname:   driver.Firstname,
		Lastname:    driver.Lastname,
		DateOfBirth: driver.DateOfBirth.AsTime(),
		Gender:      driver.Gender,
		Username:    driver.Username,
	}, nil
}

// Login ...
func (ds *DriverService) Login(username, password, carID string) (*string, error) {
	res, err := ds.DriverRepository.LoginDriver(&proto.LoginDriverRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	return &res.DriverId, nil
}

func (ds *DriverService) Update(driverID string, firstname, lastname *string, dateOfBirth *time.Time) error {
	err := ds.DriverRepository.UpdateDriver(&proto.UpdateDriverRequest{
		DriverId:    driverID,
		Firstname:   firstname,
		Lastname:    lastname,
		DateOfBirth: utils.WrapperTime(dateOfBirth),
	})
	if err != nil {
		return err
	}
	return nil
}

func (ds *DriverService) Delete(driverID string) error {
	err := ds.DriverRepository.DeleteDriver(&proto.DeleteDriverRequest{
		DriverId: driverID,
	})
	if err != nil {
		return err
	}
	return nil
}
