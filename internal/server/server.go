package server

import (
	"github.com/voroninsa/go-todo/internal/server/handlers"
	"github.com/voroninsa/go-todo/internal/storage"
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
