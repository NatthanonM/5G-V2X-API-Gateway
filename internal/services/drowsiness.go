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
func (ds *DrowsinessService) GetDrowsinessData(from, to *time.Time, carID, username *string) ([]*models.Drowsiness, error) {
	res, err := ds.DrowsinessRepository.GetDrowsinessData(&proto.GetDrowsinessDataRequest{
		From:     utils.WrapperTime(from),
		To:       utils.WrapperTime(to),
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
func (ds *DrowsinessService) GetDrowsinessStatCalendar(year int64) ([]*models.StatCal, error) {
	res, err := ds.DrowsinessRepository.GetDrowsinessStatCalendar(&year)
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
func (ds *DrowsinessService) GetNumberOfDrowsinessTimeBar(from, to *time.Time) ([]int32, error) {
	res, err := ds.DrowsinessRepository.GetNumberOfDrowsinessTimeBar(from, to)
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
	tmpDrivers := make(map[string]*models.Driver)
	tmpCars := make(map[string]*models.Car)
	for _, drowsiness := range res.Drowsinesses {
		// TODO#1: call user service to get driver id by username
		if _, ok := tmpDrivers[drowsiness.Username]; !ok {
			driver, err := ds.DriverRepository.GetDriverByUsername(&proto.GetDriverByUsernameRequest{
				Username: drowsiness.Username,
			})
			if err != nil {
				tmpDrivers[drowsiness.Username] = &models.Driver{
					Username: drowsiness.Username,
				}
				// return nil, err
			} else {
				tmpDrivers[drowsiness.Username] = &models.Driver{
					DriverID:    driver.DriverId,
					Firstname:   driver.Firstname,
					Lastname:    driver.Lastname,
					DateOfBirth: driver.DateOfBirth.AsTime(),
					Gender:      driver.Gender,
					Username:    driver.Username,
				}
			}
		}
		// TODO#2: call data-management service to get driver id by username
		if _, ok := tmpCars[drowsiness.CarId]; !ok {
			car, err := ds.CarRepository.GetCar(&proto.GetCarRequest{
				// TODO: change to accident.CarId when carId is valid
				CarId: drowsiness.CarId,
			})
			if err != nil {
				tmpCars[drowsiness.CarId] = &models.Car{
					CarID: drowsiness.CarId,
				}
				// return nil, err
			} else {
				tmpCars[drowsiness.CarId] = &models.Car{
					CarID:                     car.CarId,
					VehicleRegistrationNumber: car.VehicleRegistrationNumber,
					CarDetail:                 car.CarDetail,
					RegisteredAt:              car.RegisteredAt.AsTime(),
				}
			}
		}
		drowsinessMapData = append(drowsinessMapData, &models.DrowsinessData{
			Detail: models.DrowsinessDetail{
				Road:   drowsiness.Road,
				Time:   drowsiness.Time.AsTime(),
				Driver: tmpDrivers[drowsiness.Username],
				Car:    tmpCars[drowsiness.CarId],
			},
			Coordinate: models.Coordinate{
				Lat: drowsiness.Latitude,
				Lng: drowsiness.Longitude,
			},
		})
	}
	return drowsinessMapData, nil
}

// GetDrowsinessStatTimebar ...
func (ds *DrowsinessService) GetDrowsinessStatTimebar(from, to *time.Time, driverUsername *string) ([]int64, error) {
	drowsinessCountByHour, err := ds.DrowsinessRepository.GetDrowsinessStatGroupByHour(&proto.GetDrowsinessStatGroupByHourRequest{
		From:           utils.WrapperTime(from),
		To:             utils.WrapperTime(to),
		DriverUsername: driverUsername,
	})
	if err != nil {
		return []int64{}, err
	}
	return drowsinessCountByHour.Drowsinesses, nil
}
