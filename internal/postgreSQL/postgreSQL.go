package postgresql

import (
	"database/sql"
	"fmt"
	"serverdb/internal/domen"
	lg "serverdb/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	_ "github.com/lib/pq"
)

type PostgreSql struct {
	dsn    string
	db     *gorm.DB
	config Config
	logger *lg.CustomLogger
}

func NewPostgreSql(dbconfigpath string, logger *lg.CustomLogger) *PostgreSql {
	cfg := initPostgreSqlConfig(dbconfigpath)
	return &PostgreSql{
		config: *cfg,
		logger: logger,
	}
}

func (postgre *PostgreSql) MustInitDB() error {
	err := postgre.createDBIfNotExists()
	if err != nil {
		postgre.logger.Error("Error on creating database: ", err)
		panic(err)
	}

	err = postgre.connectDB()
	if err != nil {
		postgre.logger.Error("Error on connecting database: ", err)
		panic(err)
	}

	err = postgre.db.AutoMigrate(&domen.User{})
	if err != nil {
		postgre.logger.Error("Error on automigrating database: ", err)
		panic(err)
	}

	postgre.logger.Info("Database was initilized")

	return nil
}

func (postgre *PostgreSql) createDBIfNotExists() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable",
		postgre.config.Host, postgre.config.Port, postgre.config.User, postgre.config.Password)
	postgre.dsn = dsn

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = '%s')", postgre.config.DBName)
	err = db.QueryRow(query).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		postgre.logger.Info("Database was not exists, creating database")
		createQuery := fmt.Sprintf("CREATE DATABASE %s", postgre.config.DBName)

		if postgre.config.Tablespace != "" {
			createQuery += fmt.Sprintf(" TABLESPACE %s", postgre.config.Tablespace)
		}

		_, err = db.Exec(createQuery)
		if err != nil {
			return err
		}
	} else {
		postgre.logger.Info("Database exists")
	}
	return nil
}

func (postgre *PostgreSql) connectDB() error {
	db, err := gorm.Open(postgres.Open(postgre.dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return err
	}

	postgre.db = db
	return nil
}
