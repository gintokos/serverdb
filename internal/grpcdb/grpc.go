package grpcdb

import (
	"context"
	generated "vpagrSite/pkg/protos/gen/v1"

	"google.golang.org/grpc"
)

type serverAPI struct {
	generated.UnimplementedDBServer
	db DB
}

func NewServerAPI(db DB) *serverAPI {
	if err := db.MustInitDB(); err != nil {
		panic(err)
	}
	return &serverAPI{
		db: db,
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
