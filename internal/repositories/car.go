package repositories

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/infrastructures/grpc"
	proto "5g-v2x-api-gateway-service/pkg/api"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
)

// CarRepository ...
type CarRepository struct {
	config *config.Config
	GRPC   grpc.GRPC
}

// NewCarRepository ...
func NewCarRepository(c *config.Config, g grpc.GRPC) *CarRepository {
	return &CarRepository{
		config: c,
		GRPC:   g,
	}
}

func (cr *CarRepository) RegisterNewCar(req *proto.RegisterNewCarRequest) (*proto.RegisterNewCarResponse, error) {
	//	Connect to gRPC service
	cc := cr.GRPC.ClientConn(cr.config.DataManagementServiceConnection)
	defer cc.Close()

	res, err := proto.NewCarServiceClient(cc).RegisterNewCar(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (cr *CarRepository) GetCarList() (*proto.GetCarListResponse, error) {
	//	Connect to gRPC service
	cc := cr.GRPC.ClientConn(cr.config.DataManagementServiceConnection)
	defer cc.Close()

	res, err := proto.NewCarServiceClient(cc).GetCarList(context.Background(), new(empty.Empty))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (cr *CarRepository) GetCar(req *proto.GetCarRequest) (*proto.Car, error) {
	//	Connect to gRPC service
	cc := cr.GRPC.ClientConn(cr.config.DataManagementServiceConnection)
	defer cc.Close()

	res, err := proto.NewCarServiceClient(cc).GetCar(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (cr *CarRepository) UpdateCar(req *proto.UpdateCarRequest) error {
	cc := cr.GRPC.ClientConn(cr.config.DataManagementServiceConnection)
	defer cc.Close()

	_, err := proto.NewCarServiceClient(cc).UpdateCar(context.Background(), req)

	if err != nil {
		return err
	}

	return nil
}
