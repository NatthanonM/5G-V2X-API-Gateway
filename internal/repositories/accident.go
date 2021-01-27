package repositories

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/infrastructures/grpc"
	proto "5g-v2x-api-gateway-service/pkg/api"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

// AccidentRepository ...
type AccidentRepository struct {
	config *config.Config
	GRPC   grpc.GRPC
}

// NewAccidentRepository ...
func NewAccidentRepository(c *config.Config, g grpc.GRPC) *AccidentRepository {
	return &AccidentRepository{
		config: c,
		GRPC:   g,
	}
}

func (r *AccidentRepository) GetDailyAccidentMap(req *proto.GetHourlyAccidentOfCurrentDayRequest) (*proto.GetHourlyAccidentOfCurrentDayResponse, error) {
	//	Connect to gRPC service
	cc := r.GRPC.ClientConn(r.config.DataManagementServiceConnection)
	defer cc.Close()

	res, err := proto.NewAccidentServiceClient(cc).GetHourlyAccidentOfCurrentDay(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *AccidentRepository) GetAccidentStatCalendar() (*proto.GetNumberOfAccidentCurrentYearResponse, error) {
	//	Connect to gRPC service
	cc := r.GRPC.ClientConn(r.config.DataManagementServiceConnection)
	defer cc.Close()
	
	res, err := proto.NewAccidentServiceClient(cc).GetNumberOfAccidentCurrentYear(context.Background(),&empty.Empty{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

