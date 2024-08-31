package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/voroninsa/go-todo/storage"
	"github.com/voroninsa/go-todo/utils/common"
)

type dueHandlersGetter interface {
	GetTasksByDueDateHandler(w http.ResponseWriter, r *http.Request)
}

type dueHandlers struct {
	store *storage.TaskStore
}

func NewDueHandlers(storage *storage.TaskStore) dueHandlersGetter {
	return &dueHandlers{
		store: storage,
	}
}

func (ss *dueHandlers) GetTasksByDueDateHandler(w http.ResponseWriter, r *http.Request) {
	year, month, day, err := common.UrlToDate(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tasks, err := json.Marshal(ss.store.GetTasksByDueDate(year, month, day))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(tasks)
}