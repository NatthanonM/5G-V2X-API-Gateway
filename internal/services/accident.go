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
	config             *config.Config
}

// NewAccidentService ...
func NewAccidentService(repo *repositories.AccidentRepository, cf *config.Config) *AccidentService {
	return &AccidentService{
		AccidentRepository: repo,
		config:             cf,
	}
}

func (as *AccidentService) GetAccidentCar(from, to time.Time) ([]*models.Coordinate, error) {
	res, err := as.AccidentRepository.GetAccidentCar(&proto.GetAccidentDataRequest{
		From: utils.WrapperTime(&from),
		To:   utils.WrapperTime(&to),
	})

	if err != nil {
		return nil, err
	}
	coordinates := []*models.Coordinate{}
	for _, accident := range res.Accidents {
		coordinates = append(coordinates, &models.Coordinate{
			Lat: accident.Latitude,
			Lng: accident.Longitude,
		},
		)
	}
	return coordinates, nil
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
