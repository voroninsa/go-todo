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
	congif := config.NewConfig(flags.ConfigPath, logger)
	storage := storage.NewStorage(*congif)

	serverParams := http.ServerParams{
		Logger:  logger,
		Storage: &storage,
		Config:  congif,
	}
	server := http.NewServer(serverParams)
	server.Run()
}
