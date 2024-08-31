package handlers

import "github.com/voroninsa/go-todo/storage"

type handlersGetter interface {
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

func NewHandlers(storage *storage.TaskStore) handlersGetter {
	return &Handlers{
		NewTaskHandlers(storage).(*taskHandlers),
		NewTagHandlers(storage).(*tagHandlers),
		NewDueHandlers(storage).(*dueHandlers),
		NewAuthHandlers(storage).(*authHandlers),
	}
}
