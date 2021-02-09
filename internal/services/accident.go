package services

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/models"
	"5g-v2x-api-gateway-service/internal/repositories"
	"5g-v2x-api-gateway-service/internal/utils"
	proto "5g-v2x-api-gateway-service/pkg/api"
	"time"
)

// AccidentService ...
type AccidentService struct {
	AccidentRepository *repositories.AccidentRepository
	DriverRepository   *repositories.DriverRepository
	CarRepository      *repositories.CarRepository
	config             *config.Config
}

// NewAccidentService ...
func NewAccidentService(repo *repositories.AccidentRepository, driverRepo *repositories.DriverRepository, carRepo *repositories.CarRepository, cf *config.Config) *AccidentService {
	return &AccidentService{
		AccidentRepository: repo,
		DriverRepository:   driverRepo,
		CarRepository:      carRepo,
		config:             cf,
	}
}

func (as *AccidentService) GetAccidentCar(from, to time.Time) ([]*models.Accident, error) {
	res, err := as.AccidentRepository.GetAccidentCar(&proto.GetAccidentDataRequest{
		From: utils.WrapperTime(&from),
		To:   utils.WrapperTime(&to),
	})

	if err != nil {
		return nil, err
	}
	accidents := []*models.Accident{}
	for _, accident := range res.Accidents {
		accidents = append(accidents, &models.Accident{
			Detail: models.AccidentDetail{
				Time: accident.Time.AsTime(),
			},
			Coordinate: models.Coordinate{
				Lat: accident.Latitude,
				Lng: accident.Longitude,
			},
		})
	}
	return accidents, nil
}

func (as *AccidentService) GetDailyAccidentMap(hour int32) ([]*models.Accident, error) {
	res, err := as.AccidentRepository.GetDailyAccidentMap(&proto.GetHourlyAccidentOfCurrentDayRequest{
		Hour: hour,
	})
	if err != nil {
		return nil, err
	}
	accidents := []*models.Accident{}
	for _, accident := range res.Accidents {
		accidents = append(accidents, &models.Accident{
			Detail: models.AccidentDetail{
				Time: accident.Time.AsTime(),
			},
			Coordinate: models.Coordinate{
				Lat: accident.Latitude,
				Lng: accident.Longitude,
			},
		})
	}
	return accidents, nil
}

func (as *AccidentService) GetDailyAuthAccidentMap(hour int32) ([]*models.Accident, error) {
	res, err := as.AccidentRepository.GetDailyAccidentMap(&proto.GetHourlyAccidentOfCurrentDayRequest{
		Hour: hour,
	})
	if err != nil {
		return nil, err
	}
	accidents := []*models.Accident{}
	for _, accident := range res.Accidents {
		// TODO#1: call user service to get driver id by username
		driver, err := as.DriverRepository.GetDriverByUsername(&proto.GetDriverByUsernameRequest{
			Username: accident.Username,
		})
		if err != nil {
			return nil, err
		}
		// TODO#2: call data-management service to get driver id by username
		car, err := as.CarRepository.GetCarByID(&proto.GetCarByIDRequest{
			// TODO: change to accident.CarId when carId is valid
			CarId: "83e9b831-53f2-4e22-a4b7-039d59c69d62",
		})
		if err != nil {
			return nil, err
		}
		accidents = append(accidents, &models.Accident{
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
	return accidents, nil
}
