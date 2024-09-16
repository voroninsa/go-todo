package http

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/voroninsa/go-todo/http/handlers"
)

type routes struct {
	handlers handlers.HandlersGetter
}

func NewRouter(handlers handlers.HandlersGetter) *http.ServeMux {
	r := &routes{
		handlers: handlers,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/task/", r.taskHandlers)
	mux.HandleFunc("/due/", r.dueHandlers)
	mux.HandleFunc("/tag/", r.tagHandlers)
	mux.HandleFunc("/auth/", r.authHandler)

	fs := http.FileServer(http.Dir("./web/build"))
	mux.Handle("/", fs)

	return mux
}

func (rs *routes) taskHandlers(w http.ResponseWriter, r *http.Request) {
	// get/post/delete all tasks
	if r.URL.Path == "/task/" {
		switch r.Method {
		case http.MethodGet:
			rs.handlers.GetAllTasksHandler(w, r)
		case http.MethodPost:
			rs.handlers.PostTaskHandler(w, r)
		case http.MethodDelete:
			rs.handlers.DeleteAllTasksHandler(w, r)
		case http.MethodOptions:
			rs.handlers.OptionsHandler(w, r)
		default:
			http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
			return
		}
	} else {
		// get/post/delete task by id
		str := strings.Split(r.URL.Path, "/task/")[1]
		id, err := strconv.Atoi(str)
		if err != nil {
			http.Error(w, "incorrect id", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			rs.handlers.GetTaskHandler(w, r, id)
		case http.MethodDelete:
			rs.handlers.DeleteTaskHandler(w, r, id)
		case http.MethodPatch:
			rs.handlers.PatchTaskHandler(w, r, id)
		default:
			http.Error(w, "invalid http method (request with id)", http.StatusMethodNotAllowed)
			return
		}

	}
}

func (rs *routes) dueHandlers(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/due/" {
		switch r.Method {
		case http.MethodGet:
			rs.handlers.GetTasksByDueDateHandler(w, r)
		default:
			http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
			return
		}

	} else {
		http.Error(w, "please enter date", http.StatusBadRequest)
		return
	}

}

func (rs *routes) tagHandlers(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/tag/" {
		switch r.Method {
		case http.MethodGet:
			rs.handlers.GetTasksByTagHandler(w, r)
		default:
			http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
			return
		}

	} else {
		http.Error(w, "please enter tag", http.StatusBadRequest)
		return
	}
}

func (rs *routes) authHandler(w http.ResponseWriter, r *http.Request) {}
