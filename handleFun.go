package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "route1")
}
func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "route2")
}
func main() {
	http.HandleFunc("/route1/", d)
	http.HandleFunc("/route2/", c)
	http.ListenAndServe(":8080", nil)
}
