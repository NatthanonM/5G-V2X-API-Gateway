package repositories

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/infrastructures/grpc"
	proto "5g-v2x-api-gateway-service/pkg/api"
	"context"
)

// DriverRepository ...
type DriverRepository struct {
	config *config.Config
	GRPC   grpc.GRPC
}

// NewDriverRepository ...
func NewDriverRepository(c *config.Config, g grpc.GRPC) *DriverRepository {
	return &DriverRepository{
		config: c,
		GRPC:   g,
	}
}

func (dr *DriverRepository) AddNewDriver(req *proto.AddNewDriverRequest) (*proto.AddNewDriverReponse, error) {
	//	Connect to gRPC service
	cc := dr.GRPC.ClientConn(dr.config.UserServiceConnection)
	defer cc.Close()

	res, err := proto.NewDriverServiceClient(cc).AddNewDriver(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (dr *DriverRepository) GetDriverByUsername(req *proto.GetDriverByUsernameRequest) (*proto.GetDriverByUsernameResponse, error) {
	cc := dr.GRPC.ClientConn(dr.config.UserServiceConnection)
	defer cc.Close()

	res, err := proto.NewDriverServiceClient(cc).GetDriverByUsername(context.Background(), req)

	if err != nil {
		return nil, err
	}

	return res, nil
}
