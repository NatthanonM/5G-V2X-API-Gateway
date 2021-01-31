package repositories

import (
	"5g-v2x-api-gateway-service/internal/config"
	"5g-v2x-api-gateway-service/internal/infrastructures/grpc"
	proto "5g-v2x-api-gateway-service/pkg/api"
	"context"
)

// AdminRepository ...
type AdminRepository struct {
	config *config.Config
	GRPC   grpc.GRPC
}

// NewAdminRepository ...
func NewAdminRepository(c *config.Config, g grpc.GRPC) *AdminRepository {
	return &AdminRepository{
		config: c,
		GRPC:   g,
	}
}

func (ar *AdminRepository) Register(req *proto.RegisterAdminRequest) error {
	//	Connect to gRPC service
	cc := ar.GRPC.ClientConn(ar.config.UserServiceConnection)
	defer cc.Close()

	_, err := proto.NewAdminServiceClient(cc).RegisterAdmin(context.Background(), req)
	if err != nil {
		return err
	}

	return nil
}
