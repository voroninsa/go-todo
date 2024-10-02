package main

import (
	"github.com/voroninsa/go-todo/config"
	"github.com/voroninsa/go-todo/http"
	"github.com/voroninsa/go-todo/storage"
	"github.com/voroninsa/go-todo/utils/flags"
	"github.com/voroninsa/go-todo/utils/logger"
)

func main() {
	logger := logger.NewLogger()
	flags := flags.ParseFlags()
	config := config.NewConfig(flags.ConfigPath, logger)
	storage := storage.BackendFactory(config, logger)

	serverParams := http.ServerParams{
		Logger:  logger,
		Storage: &storage,
		Config:  config,
	}
	server := http.NewServer(serverParams)
	server.Run()
}
