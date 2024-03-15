package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (ss *Handlers) GetTasksByTagHandler(w http.ResponseWriter, r *http.Request) {
	tag := strings.Split(r.URL.Path, "/tag/")[1]

	tasks, err := json.Marshal(ss.store.GetTasksByTag(tag))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(tasks)
}
