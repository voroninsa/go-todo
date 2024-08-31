package http

import (
	"net/http"
	"os"

	"github.com/voroninsa/go-todo/http/handlers"
	"github.com/voroninsa/go-todo/storage"

	"go.uber.org/zap"
)

type ServerRunner interface {
	Run()
}

type serverImpl struct {
	store    *storage.TaskStore
	handlers handlers.Handlers
	mux      *http.ServeMux
	logger   *zap.Logger
}

type ServerParams struct {
	Storage *storage.TaskStore
	Logger  *zap.Logger
}

func NewServer(params ServerParams) ServerRunner {
	handlers := handlers.NewHandlers(params.Storage).(*handlers.Handlers)
	mux := NewRouter(*handlers)
	return &serverImpl{
		store:    params.Storage,
		handlers: *handlers,
		mux:      mux,
		logger:   params.Logger,
	}
}

func (s *serverImpl) Run() {
	s.logger.Info("♥ server start listening at port :8080 ♥")
	err := http.ListenAndServe(":8080", s.mux)
	if err != nil {
		s.logger.Sugar().Fatalf("fatal server error: ", err.Error())
		os.Exit(1)
	}
}
