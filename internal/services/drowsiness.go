package services

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/models"
	"5g-v2x-api-gateway-service/internal/repositories"
	"5g-v2x-api-gateway-service/internal/utils"
	proto "5g-v2x-api-gateway-service/pkg/api"
	"time"
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

// GetDailyDrowsinessHeatmap ...
func (ds *DrowsinessService) GetDailyDrowsinessHeatmap(from, to *time.Time) ([]*models.DrowsinessData, error) {
	res, err := ds.DrowsinessRepository.GetDrowsinessData(&proto.GetDrowsinessDataRequest{
		From: utils.WrapperTime(from),
		To:   utils.WrapperTime(to),
	})
	if err != nil {
		return nil, err
	}
	publicDrowsinessData := []*models.DrowsinessData{}
	for _, drowsiness := range res.Drowsinesses {
		publicDrowsinessData = append(publicDrowsinessData, &models.DrowsinessData{
			Detail: models.DrowsinessDetail{
				Road: drowsiness.Road,
				Time: drowsiness.Time.AsTime(),
			},
			Coordinate: models.Coordinate{
				Lat: drowsiness.Latitude,
				Lng: drowsiness.Longitude,
			},
		})
	}
	return publicDrowsinessData, nil
}

// GetDrowsinessData ...
func (ds *DrowsinessService) GetDrowsinessData(carID, username *string) ([]*models.Drowsiness, error) {
	res, err := ds.DrowsinessRepository.GetDrowsinessData(&proto.GetDrowsinessDataRequest{
		CarId:    carID,
		Username: username,
	})
	if err != nil {
		return nil, err
	}
	drowsinessData := []*models.Drowsiness{}
	for _, drowsiness := range res.Drowsinesses {
		drowsinessData = append(drowsinessData, &models.Drowsiness{
			CarID:        drowsiness.CarId,
			Username:     drowsiness.Username,
			Time:         drowsiness.Time.AsTime(),
			ResponseTime: drowsiness.ResponseTime,
			WorkingHour:  drowsiness.WorkingHour,
			Latitude:     drowsiness.Latitude,
			Longitude:    drowsiness.Longitude,
			Road:         drowsiness.Road,
		})
	}
	return drowsinessData, nil
}

// GetDrowsinessStatCalendar ...
func (ds *DrowsinessService) GetDrowsinessStatCalendar() ([]*models.StatCal, error) {
	res, err := ds.DrowsinessRepository.GetDrowsinessStatCalendar()
	if err != nil {
		return nil, err
	}
	drowsinesses := []*models.StatCal{}
	for _, drowsiness := range res.Drowsinesss {
		drowsinesses = append(drowsinesses, &models.StatCal{
			Name: drowsiness.Name,
			Data: drowsiness.Data,
		})
	}
	return drowsinesses, nil
}

// GetNumberOfDrowsinessTimeBar ...
func (ds *DrowsinessService) GetNumberOfDrowsinessTimeBar() ([]int32, error) {
	res, err := ds.DrowsinessRepository.GetNumberOfDrowsinessTimeBar()
	if err != nil {
		return nil, err
	}
	var drowsinesss []int32 = res.Drowsinesss

	return drowsinesss, nil
}

// GetDailyAuthDrowsinessHeatmap ...
func (ds *DrowsinessService) GetDailyAuthDrowsinessHeatmap(from, to *time.Time) ([]*models.DrowsinessData, error) {
	res, err := ds.DrowsinessRepository.GetDrowsinessData(&proto.GetDrowsinessDataRequest{
		From: utils.WrapperTime(from),
		To:   utils.WrapperTime(to),
	})
	if err != nil {
		return nil, err
	}
	drowsinessMapData := []*models.DrowsinessData{}
	for _, drowsiness := range res.Drowsinesses {
		// TODO#1: call user service to get driver id by username
		driver, err := ds.DriverRepository.GetDriverByUsername(&proto.GetDriverByUsernameRequest{
			Username: drowsiness.Username,
		})
		if err != nil {
			return nil, err
		}
		// TODO#2: call data-management service to get driver id by username
		car, err := ds.CarRepository.GetCar(&proto.GetCarRequest{
			// TODO: change to drowsiness.CarId when carId is valid
			CarId: drowsiness.CarId,
		})
		if err != nil {
			return nil, err
		}
		drowsinessMapData = append(drowsinessMapData, &models.DrowsinessData{
			Detail: models.DrowsinessDetail{
				Road: drowsiness.Road,
				Time: drowsiness.Time.AsTime(),
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
				Lat: drowsiness.Latitude,
				Lng: drowsiness.Longitude,
			},
		})
	}
	return drowsinessMapData, nil
}
