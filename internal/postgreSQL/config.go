package postgresql

import (
	"encoding/json"
	"os"
)

type Config struct {
    ConfigPostgresSql ConfigPostgresSql `json:"ConfigPostgresSql"`
}

type ConfigPostgresSql struct {
	DBName         string `json:"dbName"`
	User           string `json:"user"`
	Password       string `json:"password"`
	Host           string `json:"host"`
	Port           string `json:"port"`
	Tablespace     string `json:"tablespace"`
	TablespacePath string `json:"tablespacepath"`
}

func initPostgreSqlConfig(configPath string) *ConfigPostgresSql {
	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	return &config.ConfigPostgresSql
}
