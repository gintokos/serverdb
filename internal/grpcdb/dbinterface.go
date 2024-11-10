package grpcdb

import "github.com/gintokos/serverdb/internal/domen"

type DB interface {
	MustInitDB() error
	CreateUserRecord(id int64) error
	GetUserRecord(id int64) (domen.User, error)
}
