package handlers

import "github.com/voroninsa/go-todo/storage"

type authHandlersGetter interface{}

type authHandlers struct {
	store storage.Backend
}

func NewAuthHandlers(storage storage.Backend) authHandlersGetter {
	return &authHandlers{
		store: storage,
	}
}
