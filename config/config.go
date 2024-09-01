package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"go.uber.org/zap"
)

type Config struct {
	Port int `json:"port"`

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

	return &config, nil
}
