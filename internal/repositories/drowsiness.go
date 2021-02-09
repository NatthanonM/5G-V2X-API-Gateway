package repositories

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/infrastructures/grpc"
	proto "5g-v2x-api-gateway-service/pkg/api"
	"context"
)

// DrowsinessRepository ...
type DrowsinessRepository struct {
	config *config.Config
	GRPC   grpc.GRPC
}

// NewDrowsinessRepository ...
func NewDrowsinessRepository(c *config.Config, g grpc.GRPC) *DrowsinessRepository {
	return &DrowsinessRepository{
		config: c,
		GRPC:   g,
	}
}

func (r *DrowsinessRepository) GetDailyDrowsinessHeatmap(req *proto.GetHourlyDrowsinessOfCurrentDayRequest) (*proto.GetHourlyDrowsinessOfCurrentDayResponse, error) {
	//	Connect to gRPC service
	cc := r.GRPC.ClientConn(r.config.DataManagementServiceConnection)
	defer cc.Close()

	res, err := proto.NewDrowsinessServiceClient(cc).GetHourlyDrowsinessOfCurrentDay(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *DrowsinessRepository) GetDailyAuthDrowsinessHeatmap(req *proto.GetHourlyDrowsinessOfCurrentDayRequest) (*proto.GetHourlyDrowsinessOfCurrentDayResponse, error) {
	//	Connect to gRPC service
	cc := r.GRPC.ClientConn(r.config.DataManagementServiceConnection)
	defer cc.Close()

	res, err := proto.NewDrowsinessServiceClient(cc).GetHourlyDrowsinessOfCurrentDay(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
