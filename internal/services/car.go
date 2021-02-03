package services

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/models"
	"5g-v2x-api-gateway-service/internal/repositories"
	proto "5g-v2x-api-gateway-service/pkg/api"
)

// CarService ...
type CarService struct {
	CarRepository *repositories.CarRepository
	config        *config.Config
}

// NewCarService ...
func NewCarService(repo *repositories.CarRepository, cf *config.Config) *CarService {
	return &CarService{
		CarRepository: repo,
		config:        cf,
	}
}

func (cs *CarService) RegisterNewCar(carDetail, vehicleRegistrationNumber string) (*string, error) {
	request := proto.RegisterNewCarRequest{
		CarDetail:                 carDetail,
		VehicleRegistrationNumber: vehicleRegistrationNumber,
	}

	res, err := cs.CarRepository.RegisterNewCar(&request)
	if err != nil {
		return nil, err
	}
	return &res.CarId, nil
}

func (cs *CarService) GetCarList() ([]*models.Car, error) {

	res, err := cs.CarRepository.GetCarList()
	if err != nil {
		return nil, err
	}

	carList := []*models.Car{}
	for _, car := range res.CarList {
		carList = append(carList, &models.Car{
			CarID:                     car.CarId,
			VehicleRegistrationNumber: car.VehicleRegistrationNumber,
			CarDetail:                 car.CarDetail,
			RegisteredAt:              car.RegisteredAt.AsTime(),
		})
	}
	return carList, nil
}