package grpcdb

type DB interface {
	MustInitDB() error
}