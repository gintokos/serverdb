package grpcdb

type DB interface {
	MustInitDB() error
	CreateUserRecord(id int64) error
	GetUserRecord(id int64) error
}
