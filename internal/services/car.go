package services

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/models"
	"5g-v2x-api-gateway-service/internal/repositories"
	proto "5g-v2x-api-gateway-service/pkg/api"
	"time"

	"github.com/golang/protobuf/ptypes"
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

// RegisterNewCar ...
func (cs *CarService) RegisterNewCar(carDetail, vehicleRegistrationNumber string, mfgAt time.Time) (*string, error) {
	protoTime, err := ptypes.TimestampProto(mfgAt)

	request := proto.RegisterNewCarRequest{
		CarDetail:                 carDetail,
		VehicleRegistrationNumber: vehicleRegistrationNumber,
		MfgAt:                     protoTime,
	}

	res, err := cs.CarRepository.RegisterNewCar(&request)
	if err != nil {
		return nil, err
	}
	return &res.CarId, nil
}

// GetCarList ...
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
			MfgAt:                     car.MfgAt.AsTime(),
		})
	}
	return carList, nil
}

// GetCar ...
func (cs *CarService) GetCar(carID string) (*models.Car, error) {
	car, err := cs.CarRepository.GetCar(&proto.GetCarRequest{
		CarId: carID,
	})
	if err != nil {
		return nil, err
	}

	return &models.Car{
		CarID:                     car.CarId,
		VehicleRegistrationNumber: car.VehicleRegistrationNumber,
		CarDetail:                 car.CarDetail,
		RegisteredAt:              car.RegisteredAt.AsTime(),
		MfgAt:                     car.MfgAt.AsTime(),
	}, nil
}

// Update ...
func (cs *CarService) Update(carID string, carDetail, vehicleRegistrationNumber *string) error {
	err := cs.CarRepository.UpdateCar(&proto.UpdateCarRequest{
		CarId:                     carID,
		CarDetail:                 carDetail,
		VehicleRegistrationNumber: vehicleRegistrationNumber,
	})
	if err != nil {
		return err
	}
	return nil
}
