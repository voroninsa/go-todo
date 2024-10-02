package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"go.uber.org/zap"
)

type Config struct {
	// Server's port
	Port int `json:"port"`

	// Storage type (postgres, inmemory)
	StorageType string `json:"storage_type"`

	// Database's preferences
	Db_Host string `json:"db_host"`
	Db_Port int    `json:"db_port"`
	Db_Name string `json:"db_name"`
	Db_User string `json:"db_user"`
	Db_Pass string `json:"db_pass"`
}

func NewConfig(configPath string, logger *zap.Logger) *Config {
	config, err := getConfig(configPath)
	if err != nil {
		logger.Fatal("error getting config", zap.Error(err))
		os.Exit(1)
	}

	return config
}

func getConfig(configPath string) (*Config, error) {
	// Чтение конфигурационного файла
	configFile, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %w", err)
	}
	defer configFile.Close()

	// Чтение содержимого файла
	byteValue, _ := io.ReadAll(configFile)

	// Парсинг JSON в структуру Config
	var config Config
	json.Unmarshal(byteValue, &config)

	// Проверка типа хранилища
	if err := storageTypeParser(config.StorageType); err != nil {
		return nil, fmt.Errorf("storage type parser error: %w", err)
	}

	return &config, nil
}

func storageTypeParser(storageType string) error {
	switch storageType {
	case "postgres":
		return nil
	case "inmemory":
		return nil
	default:
		return fmt.Errorf("unknown storage type: %s", storageType)
	}
}

func (c *Config) IsPostgres() bool {
	return c.StorageType == "postgres"
}
