package db

import (
	"net"
	"vpagrSite/internal/data/db/grpcdb"
	postgresql "vpagrSite/internal/data/db/postgreSQL"

	"google.golang.org/grpc"
)

const (
	PORT = ":8081"
)

func MustStartDBserver() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}

	db := postgresql.NewPostgreSql()
	apiserver := grpcdb.NewServerAPI(db)

	grpcServer := grpc.NewServer()
	grpcdb.Register(grpcServer, apiserver)

	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}

