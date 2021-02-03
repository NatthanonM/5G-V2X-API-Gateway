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
	config               *config.Config
}

// NewDrowsinessService ...
func NewDrowsinessService(repo *repositories.DrowsinessRepository, cf *config.Config) *DrowsinessService {
	return &DrowsinessService{
		DrowsinessRepository: repo,
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
			Coordinate: models.AccidentCoordinate{
				Lat: accident.Latitude,
				Lng: accident.Longitude,
			},
		})
	}
	return drowsinessMapData, nil
}

func (as *DrowsinessService) GetDrowsinessStatCalendar() ([]*models.StatCal, error) {
	res, err := as.DrowsinessRepository.GetDrowsinessStatCalendar()
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

func (as *DrowsinessService) GetNumberOfDrowsinessTimeBar() ([]int32, error) {
	res, err := as.DrowsinessRepository.GetNumberOfDrowsinessTimeBar()
	if err != nil {
		return nil, err
	}
	var drowsinesss []int32 = res.Drowsinesss

	return drowsinesss, nil
}