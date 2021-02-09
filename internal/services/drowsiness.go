package services

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/models"
	"5g-v2x-api-gateway-service/internal/repositories"
	proto "5g-v2x-api-gateway-service/pkg/api"
)

// DrowsinessService ...
type DrowsinessService struct {
	DrowsinessRepository *repositories.DrowsinessRepository
	DriverRepository     *repositories.DriverRepository
	CarRepository        *repositories.CarRepository
	config               *config.Config
}

// NewDrowsinessService ...
func NewDrowsinessService(repo *repositories.DrowsinessRepository, driverRepo *repositories.DriverRepository, carRepo *repositories.CarRepository, cf *config.Config) *DrowsinessService {
	return &DrowsinessService{
		DrowsinessRepository: repo,
		DriverRepository:     driverRepo,
		CarRepository:        carRepo,
		config:               cf,
	}
}

func (ds *DrowsinessService) GetDailyDrowsinessHeatmap(hour int32) ([]*models.Drowsiness, error) {
	res, err := ds.DrowsinessRepository.GetDailyDrowsinessHeatmap(&proto.GetHourlyDrowsinessOfCurrentDayRequest{
		Hour: hour,
	})
	if err != nil {
		return nil, err
	}
	drowsinessMapData := []*models.Drowsiness{}
	for _, accident := range res.Drowsinesses {
		drowsinessMapData = append(drowsinessMapData, &models.Drowsiness{
			Detail: models.AccidentDetail{
				Time: accident.Time.AsTime(),
			},
			Coordinate: models.Coordinate{
				Lat: accident.Latitude,
				Lng: accident.Longitude,
			},
		})
	}
	return drowsinessMapData, nil
}
func (ds *DrowsinessService) GetDailyAuthDrowsinessHeatmap(hour int32) ([]*models.Drowsiness, error) {
	res, err := ds.DrowsinessRepository.GetDailyAuthDrowsinessHeatmap(&proto.GetHourlyDrowsinessOfCurrentDayRequest{
		Hour: hour,
	})
	if err != nil {
		return nil, err
	}
	drowsinessMapData := []*models.Drowsiness{}
	for _, accident := range res.Drowsinesses {
		// TODO#1: call user service to get driver id by username
		driver, err := ds.DriverRepository.GetDriverByUsername(&proto.GetDriverByUsernameRequest{
			Username: accident.Username,
		})
		if err != nil {
			return nil, err
		}
		// TODO#2: call data-management service to get driver id by username
		car, err := ds.CarRepository.GetCarByID(&proto.GetCarByIDRequest{
			// TODO: change to accident.CarId when carId is valid
			CarId: "83e9b831-53f2-4e22-a4b7-039d59c69d62",
		})
		if err != nil {
			return nil, err
		}
		drowsinessMapData = append(drowsinessMapData, &models.Drowsiness{
			Detail: models.AccidentDetail{
				Time: accident.Time.AsTime(),
				Driver: &models.Driver{
					DriverID: driver.DriverId,
				},
				Car: &models.Car{
					CarID:                     car.CarId,
					VehicleRegistrationNumber: car.VehicleRegistrationNumber,
					CarDetail:                 car.CarDetail,
					RegisteredAt:              car.RegisteredAt.AsTime(),
				},
			},
			Coordinate: models.Coordinate{
				Lat: accident.Latitude,
				Lng: accident.Longitude,
			},
		})
	}
	return drowsinessMapData, nil
}
