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

// fs := http.FileServer(http.Dir("../web/build"))
// http.Handle("/bir_order/v1/", http.StripPrefix("/bir_order/v1/", fs))
// http.HandleFunc("/bir_order/v1/echo", echo)
// // http.HandleFunc("/bir_order/v1/pack", packGet)
// // http.HandleFunc("/bir_order/v1/order", orderGet)
// // http.HandleFunc("/bir_order/v1/region", regionGet)

// err = http.ListenAndServe(":8080", nil)
// if err != nil {
// 	log.Fatal(err)
// }
