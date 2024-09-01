package handlers

import "github.com/voroninsa/go-todo/storage"

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

func NewHandlers(storage *storage.TaskStore) HandlersGetter {
	return &Handlers{
		NewTaskHandlers(storage).(*taskHandlers),
		NewTagHandlers(storage).(*tagHandlers),
		NewDueHandlers(storage).(*dueHandlers),
		NewAuthHandlers(storage).(*authHandlers),
	}
}
