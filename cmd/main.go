package main

import (
	"github.com/voroninsa/go-todo/http"
	"github.com/voroninsa/go-todo/storage"
	"github.com/voroninsa/go-todo/utils/logger"
)

func main() {
	logger := logger.NewLogger()
	storage := storage.NewStorage()

	serverParams := http.ServerParams{
		Logger:  logger,
		Storage: storage,
	}
	server := http.NewServer(serverParams)
	server.Run()
}
