package server

import (
	"go-todo/internal/server/handlers"
	"go-todo/internal/storage"
)

type ServerStore struct {
	Store    *storage.TaskStore
	Handlers handlers.Handlers
}

func NewStorage() *ServerStore {
	store := storage.New()
	return &ServerStore{
		Store:    store,
		Handlers: handlers.Handlers{},
	}
}
