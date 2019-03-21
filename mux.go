package main

import (
	"io"
	"net/http"
)

type handler int

func (d handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "route 1")
}

type handler2 int

func (h handler2) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "route 2")
}
func main() {
	var d handler
	var h handler2
	mux := http.NewServeMux()
	mux.Handle("/route1/", d)
	mux.Handle("/route2", h)
	http.ListenAndServe(":8080", mux)
}
