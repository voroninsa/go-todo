package server

import (
	"net/http"
	"strconv"
	"strings"
)

func (ss *ServerStore) TaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/task/" {
		switch r.Method {
		case http.MethodGet:
			ss.Handlers.GetAllTasksHandler(w, r)
		case http.MethodPost:
			ss.Handlers.PostTaskHandler(w, r)
		case http.MethodDelete:
			ss.Handlers.DeleteAllTasksHandler(w, r)
		case http.MethodOptions:
			ss.Handlers.OptionsHandler(w, r)
		default:
			http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
			return
		}
	} else {
		str := strings.Split(r.URL.Path, "/task/")[1]
		id, err := strconv.Atoi(str)
		if err != nil {
			http.Error(w, "incorrect id", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			ss.Handlers.GetTaskHandler(w, r, id)
		case http.MethodDelete:
			ss.Handlers.DeleteTaskHandler(w, r, id)
		case http.MethodPatch:
			ss.Handlers.PatchTaskHandler(w, r, id)
		default:
			http.Error(w, "invalid http method (request with id)", http.StatusMethodNotAllowed)
			return
		}

	}
}

func (ss *ServerStore) DueHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/due/" {
		switch r.Method {
		case http.MethodGet:
			ss.Handlers.GetTasksByDueDateHandler(w, r)
		default:
			http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
			return
		}

	} else {
		http.Error(w, "please enter date", http.StatusBadRequest)
		return
	}

}

func (ss *ServerStore) TagHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/tag/" {
		switch r.Method {
		case http.MethodGet:
			ss.Handlers.GetTasksByTagHandler(w, r)
		default:
			http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
			return
		}

	} else {
		http.Error(w, "please enter tag", http.StatusBadRequest)
		return
	}
}
