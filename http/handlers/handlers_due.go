package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/voroninsa/go-todo/storage"
	"github.com/voroninsa/go-todo/utils/common"
	"github.com/voroninsa/go-todo/utils/dto"
)

type dueHandlersGetter interface {
	GetTasksByDueDateHandler(w http.ResponseWriter, r *http.Request)
}

type dueHandlers struct {
	store storage.Backend
}

func NewDueHandlers(storage storage.Backend) dueHandlersGetter {
	return &dueHandlers{
		store: storage,
	}
}

func (d *dueHandlers) GetTasksByDueDateHandler(w http.ResponseWriter, r *http.Request) {
	date, err := common.UrlToDate(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	storageResp, err := d.store.Read(dto.StorageRequest{
		Target: dto.RequestTargetDueDate,
		Task: dto.Task{
			Due: date,
		},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(storageResp.Tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}
