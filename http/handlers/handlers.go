package handlers

import (
	"github.com/voroninsa/go-todo/storage"
	"go.uber.org/zap"
)

const (
	errUnexpectedError = "unexpected error"
)

type HandlersGetter interface {
	taskHandlersGetter
	tagHandlersGetter
	dueHandlersGetter
	authHandlersGetter
}

type Handlers struct {
	*taskHandlers
	*tagHandlers
	*dueHandlers
	*authHandlers
}

func NewHandlers(storage storage.Backend, logger *zap.Logger) HandlersGetter {
	return &Handlers{
		NewTaskHandlers(storage, logger).(*taskHandlers),
		NewTagHandlers(storage, logger).(*tagHandlers),
		NewDueHandlers(storage, logger).(*dueHandlers),
		NewAuthHandlers(storage, logger).(*authHandlers),
	}
}
