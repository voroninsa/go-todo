package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/voroninsa/go-todo/storage"
)

type tagHandlersGetter interface {
	GetTasksByTagHandler(w http.ResponseWriter, r *http.Request)
}

type tagHandlers struct {
	store *storage.TaskStore
}

func NewTagHandlers(storage *storage.TaskStore) tagHandlersGetter {
	return &tagHandlers{
		store: storage,
	}
}

func (ss *tagHandlers) GetTasksByTagHandler(w http.ResponseWriter, r *http.Request) {
	tag := strings.Split(r.URL.Path, "/tag/")[1]

	tasks, err := json.Marshal(ss.store.GetTasksByTag(tag))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(tasks)
}
