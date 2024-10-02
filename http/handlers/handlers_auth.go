package handlers

import (
	"github.com/voroninsa/go-todo/storage"
	"go.uber.org/zap"
)

type authHandlersGetter interface{}

type authHandlers struct {
	store  storage.Backend
	logger *zap.Logger
}

func NewAuthHandlers(storage storage.Backend, logger *zap.Logger) authHandlersGetter {
	log := logger.Named("handlers_auth")

	return &authHandlers{
		store:  storage,
		logger: log,
	}
}
