package db

import (
	"net"
	"serverdb/internal/grpcdb"
	postgresql "serverdb/internal/postgreSQL"
	"serverdb/pkg/logger"

	"google.golang.org/grpc"
)

const (
	PORT = ":8081"
)

func MustStartDBserver(configpath string) {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}

	logger.MustSetupLogger()
	lg := logger.GetLogger()

	lg.Info("Logger was initilized")

	db := postgresql.NewPostgreSql(configpath, lg)

	apiserver := grpcdb.NewServerAPI(db, lg)
	lg.Info("Server was initilized")

	grpcServer := grpc.NewServer()
	grpcdb.Register(grpcServer, apiserver)

	err = grpcServer.Serve(lis)
	if err != nil {
		lg.Error("Error on serving server ", err)
		panic(err)
	}
}
