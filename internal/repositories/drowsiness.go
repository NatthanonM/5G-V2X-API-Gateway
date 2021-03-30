package repositories

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/infrastructures/grpc"
	"5g-v2x-api-gateway-service/internal/utils"
	proto "5g-v2x-api-gateway-service/pkg/api"
	"context"
	"time"
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

func (r *DrowsinessRepository) GetDrowsinessData(req *proto.GetDrowsinessDataRequest) (*proto.GetDrowsinessDataResponse, error) {
	//	Connect to gRPC service
	cc := r.GRPC.ClientConn(r.config.DataManagementServiceConnection)
	defer cc.Close()

	res, err := proto.NewDrowsinessServiceClient(cc).GetDrowsinessData(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *DrowsinessRepository) GetDrowsinessStatCalendar(year *int64) (*proto.GetNumberOfDrowsinessToCalendarResponse, error) {
	//	Connect to gRPC service
	cc := r.GRPC.ClientConn(r.config.DataManagementServiceConnection)
	defer cc.Close()

	res, err := proto.NewDrowsinessServiceClient(cc).GetNumberOfDrowsinessToCalendar(context.Background(), &proto.GetNumberOfDrowsinessToCalendarRequest{
		Year: year,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *DrowsinessRepository) GetNumberOfDrowsinessTimeBar(from, to *time.Time) (*proto.GetNumberOfDrowsinessTimeBarResponse, error) {
	//	Connect to gRPC service
	cc := r.GRPC.ClientConn(r.config.DataManagementServiceConnection)
	defer cc.Close()

	res, err := proto.NewDrowsinessServiceClient(cc).GetNumberOfDrowsinessTimeBar(context.Background(), &proto.GetNumberOfDrowsinessTimeBarRequest{
		From: utils.WrapperTime(from),
		To:   utils.WrapperTime(to),
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *DrowsinessRepository) GetDrowsinessStatGroupByHour(req *proto.GetDrowsinessStatGroupByHourRequest) (*proto.GetDrowsinessStatGroupByHourResponse, error) {
	//	Connect to gRPC service
	cc := r.GRPC.ClientConn(r.config.DataManagementServiceConnection)
	defer cc.Close()

	res, err := proto.NewDrowsinessServiceClient(cc).GetDrowsinessStatGroupByHour(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
