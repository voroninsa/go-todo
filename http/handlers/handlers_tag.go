package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/voroninsa/go-todo/storage"
	"github.com/voroninsa/go-todo/utils/dto"
)

type tagHandlersGetter interface {
	GetTasksByTagHandler(w http.ResponseWriter, r *http.Request)
}

type tagHandlers struct {
	store storage.Backend
}

func NewTagHandlers(storage storage.Backend) tagHandlersGetter {
	return &tagHandlers{
		store: storage,
	}
}

func (t *tagHandlers) GetTasksByTagHandler(w http.ResponseWriter, r *http.Request) {
	tag := strings.Split(r.URL.Path, "/tag/")[1]

	// Получение задач по тегу
	tasks, err := t.store.Read(dto.StorageRequest{
		Tag: tag,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}
