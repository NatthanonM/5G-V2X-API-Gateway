package services

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/models"
	"5g-v2x-api-gateway-service/internal/repositories"
	"5g-v2x-api-gateway-service/internal/utils"
	proto "5g-v2x-api-gateway-service/pkg/api"
	"fmt"
	"time"
)

// AccidentService ...
type AccidentService struct {
	AccidentRepository *repositories.AccidentRepository
	DriverRepository   *repositories.DriverRepository
	config             *config.Config
}

// NewAccidentService ...
func NewAccidentService(repo *repositories.AccidentRepository, driverRepo *repositories.DriverRepository, cf *config.Config) *AccidentService {
	return &AccidentService{
		AccidentRepository: repo,
		DriverRepository:   driverRepo,
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
	fmt.Println(res.Accidents)
	for _, accident := range res.Accidents {
		// TODO#1: call user service to get driver id by username
		driver, err := as.DriverRepository.GetDriverByUsername(&proto.GetDriverByUsernameRequest{
			Username: accident.Username,
		})
		if err != nil {
			return nil, err
		}
		// TODO#2: call data-management service to get driver id by username
		accidents = append(accidents, &models.Accident{
			Detail: models.AccidentDetail{
				Time: accident.Time.AsTime(),
				Driver: &models.Driver{
					DriverID: driver.DriverId,
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
