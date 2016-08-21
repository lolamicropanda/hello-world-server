package main

import (
	"net/http"
)

func initServeMux() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})
	return m
}

func main() {
	svr := &http.Server{
		Addr: ":1123",
		Handler: initServeMux(),
	}
	svr.ListenAndServe()
}
