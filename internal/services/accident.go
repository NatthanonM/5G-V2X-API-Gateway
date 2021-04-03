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

// GetAccident ...
func (as *AccidentService) GetAccident(from, to *time.Time, carID, username *string) ([]*models.Accident, error) {
	res, err := as.AccidentRepository.GetAccidentData(&proto.GetAccidentDataRequest{
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

// GetDailyAccidentMap ...
func (as *AccidentService) GetDailyAccidentMap(from, to *time.Time) ([]*models.AccidentData, error) {
	res, err := as.AccidentRepository.GetAccidentData(&proto.GetAccidentDataRequest{
		From: utils.WrapperTime(from),
		To:   utils.WrapperTime(to),
	})
	if err != nil {
		return nil, err
	}
	publicAccidentData := []*models.AccidentData{}
	for _, accident := range res.Accidents {
		publicAccidentData = append(publicAccidentData, &models.AccidentData{
			Detail: models.AccidentDetail{
				Road: accident.Road,
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

// GetAccidentStatCalendar ...
func (as *AccidentService) GetAccidentStatCalendar(year int64) ([]*models.StatCal, error) {
	res, err := as.AccidentRepository.GetAccidentStatCalendar(&year)
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

// GetNumberOfAccidentTimeBar ...
func (as *AccidentService) GetNumberOfAccidentTimeBar(from, to *time.Time) ([]int32, error) {
	res, err := as.AccidentRepository.GetNumberOfAccidentTimeBar(from, to)
	if err != nil {
		return nil, err
	}
	var accidents []int32 = res.Accidents

	return accidents, nil
}

// GetNumberOfAccidentStreet ...
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

// GetDailyAuthAccidentMap ...
func (as *AccidentService) GetDailyAuthAccidentMap(from, to *time.Time) ([]*models.AccidentData, error) {
	res, err := as.AccidentRepository.GetAccidentData(&proto.GetAccidentDataRequest{
		From: utils.WrapperTime(from),
		To:   utils.WrapperTime(to),
	})
	if err != nil {
		return nil, err
	}
	accidents := []*models.AccidentData{}
	tmpDrivers := make(map[string]*models.Driver)
	tmpCars := make(map[string]*models.Car)
	for _, accident := range res.Accidents {
		// TODO#1: call user service to get driver id by username
		if _, ok := tmpDrivers[accident.Username]; !ok {
			driver, err := as.DriverRepository.GetDriverByUsername(&proto.GetDriverByUsernameRequest{
				Username: accident.Username,
			})
			if err != nil {
				return nil, err
			}
			tmpDrivers[accident.Username] = &models.Driver{
				DriverID:    driver.DriverId,
				Firstname:   driver.Firstname,
				Lastname:    driver.Lastname,
				DateOfBirth: driver.DateOfBirth.AsTime(),
				Gender:      driver.Gender,
				Username:    driver.Username,
			}
		}
		// TODO#2: call data-management service to get driver id by username
		if _, ok := tmpCars[accident.CarId]; !ok {
			car, err := as.CarRepository.GetCar(&proto.GetCarRequest{
				// TODO: change to accident.CarId when carId is valid
				CarId: accident.CarId,
			})
			if err != nil {
				return nil, err
			}
			tmpCars[accident.CarId] = &models.Car{
				CarID:                     car.CarId,
				VehicleRegistrationNumber: car.VehicleRegistrationNumber,
				CarDetail:                 car.CarDetail,
				RegisteredAt:              car.RegisteredAt.AsTime(),
			}
		}
		accidents = append(accidents, &models.AccidentData{
			Detail: models.AccidentDetail{
				Time:   accident.Time.AsTime(),
				Road:   accident.Road,
				Driver: tmpDrivers[accident.Username],
				Car:    tmpCars[accident.CarId],
			},
			Coordinate: models.Coordinate{
				Lat: accident.Latitude,
				Lng: accident.Longitude,
			},
		})
	}
	return accidents, nil
}

// GetAccidentStatTimebar ...
func (as *AccidentService) GetAccidentStatTimebar(from, to *time.Time, driverUsername *string) ([]int64, error) {
	accidentCountByHour, err := as.AccidentRepository.GetAccidentStatGroupByHour(&proto.GetAccidentStatGroupByHourRequest{
		From:           utils.WrapperTime(from),
		To:             utils.WrapperTime(to),
		DriverUsername: driverUsername,
	})
	if err != nil {
		return []int64{}, err
	}
	return accidentCountByHour.Accidents, nil
}

func (as *AccidentService) GetStatRoadToptenYear(from, to *time.Time, n *int64) ([]models.StatRoadToptenYear, error) {
	topNAccidentRoadResponse, err := as.AccidentRepository.GetTopNRoad(&proto.GetTopNRoadRequest{
		From: utils.WrapperTime(from),
		To:   utils.WrapperTime(to),
		N:    n,
	})
	if err != nil {
		return []models.StatRoadToptenYear{}, err
	}
	var topNAccidentRoad []models.StatRoadToptenYear
	for _, v := range topNAccidentRoadResponse.TopNRoad {
		topNAccidentRoad = append(topNAccidentRoad, models.StatRoadToptenYear{
			RoadName:      v.RoadName,
			AccidentCount: v.AccidentCount,
		})
	}
	return topNAccidentRoad, nil
}
