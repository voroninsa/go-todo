package http

import (
	"fmt"
	"net/http"
	"os"

	"github.com/voroninsa/go-todo/config"
	"github.com/voroninsa/go-todo/http/handlers"
	"github.com/voroninsa/go-todo/storage"

	"go.uber.org/zap"
)

type ServerRunner interface {
	Run()
}

type serverImpl struct {
	store  *storage.TaskStore
	mux    *http.ServeMux
	config *config.Config
	logger *zap.Logger
}

type ServerParams struct {
	Storage *storage.TaskStore
	Config  *config.Config
	Logger  *zap.Logger
}

func NewServer(params ServerParams) ServerRunner {
	handlers := handlers.NewHandlers(params.Storage)
	mux := NewRouter(handlers)
	return &serverImpl{
		store:  params.Storage,
		mux:    mux,
		config: params.Config,
		logger: params.Logger,
	}
}

func (s *serverImpl) Run() {
	port := fmt.Sprintf(":%d", s.config.Port)
	s.logger.Sugar().Infof("♥ server start listening at port %s ♥", port)
	err := http.ListenAndServe(port, s.mux)
	if err != nil {
		s.logger.Sugar().Fatalf("fatal server error: ", err.Error())
		os.Exit(1)
	}
}
