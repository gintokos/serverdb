package grpcdb

import (
	"context"
	"github.com/gintokos/serverdb/pkg/logger"
	generated "github.com/gintokos/serverdb/protos/gen/v1"

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
	s.logger.Info("Get CreateUser endpoint")

	var errmsg string
	err := s.db.CreateUserRecord(req.UserId)
	if err != nil {
		s.logger.Error("Error on creating record", err)
		errmsg = err.Error()
	}

	resp := generated.CreateUserResponce{
		Result: true,
		Error:  errmsg,
	}

	return &resp, err
}

func (s *serverAPI) GetUser(ctx context.Context, req *generated.GetUserRequest) (*generated.GetUserResponce, error) {
	s.logger.Info("Get GetUser endpoint")

	var errmsg string
	user, err := s.db.GetUserRecord(req.UserId)
	if err != nil {
		s.logger.Error("Error on geting record", err)
		errmsg = err.Error()
	}

	resp := generated.GetUserResponce{
		Result:    true,
		CreatedAt: user.CreatedAt.Format("02.01.2006 15:04:05"),
		Error:     errmsg,
	}

	return &resp, err
}
