package handlers

import (
	"encoding/json"
	"go-todo/internal/common"
	"net/http"
)

func (ss *Handlers) GetTasksByDueDateHandler(w http.ResponseWriter, r *http.Request) {
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
