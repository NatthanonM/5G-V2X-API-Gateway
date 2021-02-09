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

func (as *AccidentService) GetAccident(from, to *time.Time, carID, username *string) ([]*models.Accident, error) {
	res, err := as.AccidentRepository.GetAccidentCar(&proto.GetAccidentDataRequest{
		From:     utils.WrapperTime(from),
		To:       utils.WrapperTime(to),
		CarId:    carID,
		Username: username,
	})

	if err != nil {
		return nil, err
	}
	accidents := []*models.Accident{}
	for _, accident := range res.Accidents {
		accidents = append(accidents, &models.Accident{
			CarID:     accident.CarId,
			Username:  accident.Username,
			Time:      accident.Time.AsTime(),
			Latitude:  accident.Latitude,
			Longitude: accident.Longitude,
			Road:      accident.Road,
		})
	}
	return accidents, nil
}

func (as *AccidentService) GetDailyAccidentMap(hour int32) ([]*models.PublicAccidentData, error) {
	res, err := as.AccidentRepository.GetDailyAccidentMap(&proto.GetHourlyAccidentOfCurrentDayRequest{
		Hour: hour,
	})
	if err != nil {
		return nil, err
	}
	publicAccidentData := []*models.PublicAccidentData{}
	for _, accident := range res.Accidents {
		publicAccidentData = append(publicAccidentData, &models.PublicAccidentData{
			Detail: models.AccidentDetail{
				Time: accident.Time.AsTime(),
			},
			Coordinate: models.Coordinate{
				Lat: accident.Latitude,
				Lng: accident.Longitude,
			},
		})
	}
	return publicAccidentData, nil
}

func (as *AccidentService) GetAccidentStatCalendar() ([]*models.StatCal, error) {
	res, err := as.AccidentRepository.GetAccidentStatCalendar()
	if err != nil {
		return nil, err
	}
	accidents := []*models.StatCal{}
	for _, accident := range res.Accidents {
		accidents = append(accidents, &models.StatCal{
			Name: accident.Name,
			Data: accident.Data,
		})
	}
	return accidents, nil
}

func (as *AccidentService) GetNumberOfAccidentTimeBar() ([]int32, error) {
	res, err := as.AccidentRepository.GetNumberOfAccidentTimeBar()
	if err != nil {
		return nil, err
	}
	var accidents []int32 = res.Accidents

	return accidents, nil
}

func (as *AccidentService) GetNumberOfAccidentStreet() (*models.StatPie, error) {
	res, err := as.AccidentRepository.GetNumberOfAccidentStreet()
	if err != nil {
		return nil, err
	}
	var accidents *models.StatPie
	accidents = &models.StatPie{
		Series: res.Accidents.Series,
		Labels: res.Accidents.Labels,
	}

	return accidents, nil
}
