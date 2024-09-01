package main

import (
	"github.com/voroninsa/go-todo/config"
	"github.com/voroninsa/go-todo/http"
	"github.com/voroninsa/go-todo/storage"
	"github.com/voroninsa/go-todo/utils/common"
	"github.com/voroninsa/go-todo/utils/logger"
)

func main() {
	logger := logger.NewLogger()
	storage := storage.NewStorage()
	flags := common.ParseFlags()

	configPath := flags["config"].(string)
	congif := config.NewConfig(configPath, logger)

	serverParams := http.ServerParams{
		Logger:  logger,
		Storage: storage,
		Config:  congif,
	}
	server := http.NewServer(serverParams)
	server.Run()
}
