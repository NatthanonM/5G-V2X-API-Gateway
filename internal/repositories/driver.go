package repositories

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/infrastructures/grpc"
	proto "5g-v2x-api-gateway-service/pkg/api"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
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

func (dr *DriverRepository) GetAllDriver() (*proto.GetAllDriverResponse, error) {
	//	Connect to gRPC service
	cc := dr.GRPC.ClientConn(dr.config.UserServiceConnection)
	defer cc.Close()

	res, err := proto.NewDriverServiceClient(cc).GetAllDriver(context.Background(), new(empty.Empty))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (dr *DriverRepository) GetDriver(req *proto.GetDriverRequest) (*proto.Driver, error) {
	//	Connect to gRPC service
	cc := dr.GRPC.ClientConn(dr.config.UserServiceConnection)
	defer cc.Close()

	res, err := proto.NewDriverServiceClient(cc).GetDriver(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (dr *DriverRepository) GetDriverByUsername(req *proto.GetDriverByUsernameRequest) (*proto.Driver, error) {
	cc := dr.GRPC.ClientConn(dr.config.UserServiceConnection)
	defer cc.Close()

	res, err := proto.NewDriverServiceClient(cc).GetDriverByUsername(context.Background(), req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (dr *DriverRepository) LoginDriver(req *proto.LoginDriverRequest) (*proto.LoginDriverResponse, error) {
	cc := dr.GRPC.ClientConn(dr.config.UserServiceConnection)
	defer cc.Close()

	res, err := proto.NewDriverServiceClient(cc).LoginDriver(context.Background(), req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (dr *DriverRepository) UpdateDriver(req *proto.UpdateDriverRequest) error {
	cc := dr.GRPC.ClientConn(dr.config.UserServiceConnection)
	defer cc.Close()

	_, err := proto.NewDriverServiceClient(cc).UpdateDriver(context.Background(), req)

	if err != nil {
		return err
	}

	return nil
}

func (dr *DriverRepository) DeleteDriver(req *proto.DeleteDriverRequest) error {
	cc := dr.GRPC.ClientConn(dr.config.UserServiceConnection)
	defer cc.Close()

	_, err := proto.NewDriverServiceClient(cc).DeleteDriver(context.Background(), req)

	if err != nil {
		return err
	}

	return nil
}
