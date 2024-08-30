package main

import (
	"net/http"
	"os"

	"github.com/voroninsa/go-todo/internal/server"
	"github.com/voroninsa/go-todo/internal/utils/logger"
)

func main() {
	logger := logger.NewLogger()

	storage := server.NewStorage()
	mux := http.NewServeMux()
	mux.HandleFunc("/task/", storage.TaskHandler)
	mux.HandleFunc("/due/", storage.DueHandler)
	mux.HandleFunc("/tag/", storage.TagHandler)

	fs := http.FileServer(http.Dir("./web/build"))
	mux.Handle("/", fs)

	logger.Info("♥ server start listening at port :8080 ♥")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		logger.Sugar().Fatalf("fatal server error: ", err.Error())
		os.Exit(1)
	}
}
