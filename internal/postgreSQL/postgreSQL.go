package postgresql

import (
	"database/sql"
	"fmt"
	"serverdb/internal/domen"
	lg "serverdb/pkg/logger"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	err := postgre.initializeTablespace(postgre.config.TablespacePath)
	if err != nil {
		postgre.logger.Error("Error initializing tablespace: ", err)
		panic(err)
	}

	err = postgre.createDBIfNotExists()
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

	postgre.logger.Info("Database was initialized")
	return nil
}

func (postgre *PostgreSql) initializeTablespace(storagePath string) error {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s?sslmode=disable",
		postgre.config.User,
		postgre.config.Password,
		postgre.config.Host,
		postgre.config.Port,
	)
	postgre.dsn = dsn
	postgre.logger.Info(dsn)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("error opening connection: %w", err)
	}
	defer db.Close()

	// Check if tablespace exists
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_tablespace WHERE spcname = '%s')",
		postgre.config.Tablespace)
	err = db.QueryRow(query).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking tablespace existence: %w", err)
	}

	if !exists {
		createTablespaceQuery := fmt.Sprintf("CREATE TABLESPACE %s LOCATION '%s'",
			postgre.config.Tablespace, storagePath)

		_, err = db.Exec(createTablespaceQuery)
		if err != nil {
			return fmt.Errorf("error creating tablespace: %w", err)
		}
		postgre.logger.Info("Tablespace created successfully")
	} else {
		postgre.logger.Info("Tablespace exists")
	}

	return nil
}

func (postgre *PostgreSql) createDBIfNotExists() error {
	db, err := sql.Open("postgres", postgre.dsn)
	if err != nil {
		return fmt.Errorf("error opening connection: %w", err)
	}
	defer db.Close()

	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = '%s')",
		postgre.config.DBName)
	err = db.QueryRow(query).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking database existence: %w", err)
	}

	if !exists {
		postgre.logger.Info("Database does not exist, creating database")
		createQuery := fmt.Sprintf("CREATE DATABASE %s", postgre.config.DBName)

		if postgre.config.Tablespace != "" {
			createQuery += fmt.Sprintf(" TABLESPACE %s", postgre.config.Tablespace)
		}

		_, err = db.Exec(createQuery)
		if err != nil {
			return fmt.Errorf("error creating database: %w", err)
		}
	} else {
		postgre.logger.Info("Database exists")
	}
	return nil
}

func (postgre *PostgreSql) connectDB() error {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		postgre.config.User,
		postgre.config.Password,
		postgre.config.Host,
		postgre.config.Port,
		postgre.config.DBName,
	)
	// Update connection string to include database name
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}

	postgre.db = db
	return nil
}
