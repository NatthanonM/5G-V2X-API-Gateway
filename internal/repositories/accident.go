package repositories

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/infrastructures/grpc"
	"5g-v2x-api-gateway-service/internal/utils"
	proto "5g-v2x-api-gateway-service/pkg/api"
	"context"
	"time"

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

// func (r *AccidentRepository) GetDailyAccidentMap(req *proto.GetHourlyAccidentOfCurrentDayRequest) (*proto.GetHourlyAccidentOfCurrentDayResponse, error) {
// 	//	Connect to gRPC service
// 	cc := r.GRPC.ClientConn(r.config.DataManagementServiceConnection)
// 	defer cc.Close()

// 	res, err := proto.NewAccidentServiceClient(cc).GetHourlyAccidentOfCurrentDay(context.Background(), req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return res, nil
// }

func (r *AccidentRepository) GetAccidentData(req *proto.GetAccidentDataRequest) (*proto.GetAccidentDataResponse, error) {
	//	Connect to gRPC service
	cc := r.GRPC.ClientConn(r.config.DataManagementServiceConnection)
	defer cc.Close()

	res, err := proto.NewAccidentServiceClient(cc).GetAccidentData(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *AccidentRepository) GetAccidentStatCalendar(year *int64) (*proto.GetNumberOfAccidentToCalendarResponse, error) {
	//	Connect to gRPC service
	cc := r.GRPC.ClientConn(r.config.DataManagementServiceConnection)
	defer cc.Close()

	res, err := proto.NewAccidentServiceClient(cc).GetNumberOfAccidentToCalendar(context.Background(), &proto.GetNumberOfAccidentToCalendarRequest{
		Year: year,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *AccidentRepository) GetNumberOfAccidentTimeBar(from, to *time.Time) (*proto.GetNumberOfAccidentTimeBarResponse, error) {
	//	Connect to gRPC service
	cc := r.GRPC.ClientConn(r.config.DataManagementServiceConnection)
	defer cc.Close()

	res, err := proto.NewAccidentServiceClient(cc).GetNumberOfAccidentTimeBar(context.Background(), &proto.GetNumberOfAccidentTimeBarRequest{
		From: utils.WrapperTime(from),
		To:   utils.WrapperTime(to),
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *AccidentRepository) GetNumberOfAccidentStreet() (*proto.GetNumberOfAccidentStreetResponse, error) {
	//	Connect to gRPC service
	cc := r.GRPC.ClientConn(r.config.DataManagementServiceConnection)
	defer cc.Close()

	res, err := proto.NewAccidentServiceClient(cc).GetNumberOfAccidentStreet(context.Background(), &empty.Empty{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *AccidentRepository) GetAccidentStatGroupByHour(req *proto.GetAccidentStatGroupByHourRequest) (*proto.GetAccidentStatGroupByHourResponse, error) {
	//	Connect to gRPC service
	cc := r.GRPC.ClientConn(r.config.DataManagementServiceConnection)
	defer cc.Close()

	res, err := proto.NewAccidentServiceClient(cc).GetAccidentStatGroupByHour(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *AccidentRepository) GetTopNRoad(req *proto.GetTopNRoadRequest) (*proto.GetTopNRoadResponse, error) {
	//	Connect to gRPC service
	cc := r.GRPC.ClientConn(r.config.DataManagementServiceConnection)
	defer cc.Close()

	res, err := proto.NewAccidentServiceClient(cc).GetTopNRoad(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
