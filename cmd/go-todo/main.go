package main

import (
	"fmt"
	"go-todo/internal/server"
	"log"
	"net/http"
)

func main() {
	s := server.NewStorage()
	mux := http.NewServeMux()
	mux.HandleFunc("/task/", s.TaskHandler)
	mux.HandleFunc("/due/", s.DueHandler)
	mux.HandleFunc("/tag/", s.TagHandler)

	fs := http.FileServer(http.Dir("./web/build"))
	mux.Handle("/", fs)

	fmt.Println("♥ server start listening at port :8080 ♥")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
