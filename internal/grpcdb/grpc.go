package grpcdb

import (
	"context"
	"serverdb/pkg/logger"
	generated "serverdb/protos/gen/v1"

	"google.golang.org/grpc"
)

type serverAPI struct {
	generated.UnimplementedDBServer
	logger *logger.CustomLogger
	db     DB
}

func NewServerAPI(db DB, logger *logger.CustomLogger) *serverAPI {
	if err := db.MustInitDB(); err != nil {
		panic(err)
	}
	return &serverAPI{
		logger: logger,
		db:     db,
	}
}

func Register(gRPC *grpc.Server, serverapi *serverAPI) {
	generated.RegisterDBServer(gRPC, serverapi)
}

func (s *serverAPI) CreateUser(ctx context.Context, req *generated.CreateUserRequest) (*generated.CreateUserResponce, error) {
	// uID = req.GetUserId()
	panic("impl me")

	return nil, nil
}
func (s *serverAPI) GetUser(ctx context.Context, req *generated.GetUserRequest) (*generated.GetUserResponce, error) {
	panic("impl me")
	return nil, nil
}
