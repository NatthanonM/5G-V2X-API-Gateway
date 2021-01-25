package services

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/models"
	"5g-v2x-api-gateway-service/internal/repositories"
	proto "5g-v2x-api-gateway-service/pkg/api"
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

func (as *AccidentService) GetDailyAccidentMap(hour int32) (*models.MapResponseData, error) {
	res, err := as.AccidentRepository.GetDailyAccidentMap(&proto.GetHourlyAccidentOfCurrentDayRequest{
		Hour: hour,
	})
	if err != nil {
		return nil, err
	}
	var mapResponseData models.MapResponseData
	for _, accident := range res.Accidents {
		mapResponseData.Accidents = append(mapResponseData.Accidents, models.Accident{
			Detail: models.AccidentDetail{
				Time: accident.Time.AsTime(),
			},
			Coordinate: models.AccidentCoordinate{
				Lat: accident.Latitude,
				Lng: accident.Longitude,
			},
		})
	}
	return &mapResponseData, nil
}
