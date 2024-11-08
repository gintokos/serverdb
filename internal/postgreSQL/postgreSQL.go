package postgresql

import (
	"database/sql"
	"fmt"
)

type PostgreSql struct {
}

func NewPostgreSql() *PostgreSql {
	return &PostgreSql{}
}

func (postgre PostgreSql) MustInitDB() error {
	// cfg := initPostgreSqlConfig()
	return nil
}

func createDBIfNotExists(config *Config) error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// _, err = db.Exec(fmt.Sprintf())
}
