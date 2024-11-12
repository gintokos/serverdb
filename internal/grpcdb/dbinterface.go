package grpcdb

import "github.com/gintokos/serverdb/internal/domain"

type DB interface {
	MustInitDB() error
	CreateUserRecord(id int64) error
	GetUserRecord(id int64) (domain.User, error)
}
