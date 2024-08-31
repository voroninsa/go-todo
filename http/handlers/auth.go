package handlers

import "github.com/voroninsa/go-todo/storage"

type authHandlersGetter interface{}

type authHandlers struct {
	store *storage.TaskStore
}

func NewAuthHandlers(storage *storage.TaskStore) authHandlersGetter {
	return &authHandlers{
		store: storage,
	}
}
