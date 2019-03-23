package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {
	http.HandleFunc("/", r1)
	http.ListenAndServe(":8082", nil)
	http.HandleFunc("/route1/", r2)
}
func r1(w http.ResponseWriter, req *http.Request) {
	fmt.Println("logging request method", req.Method)
}
func r2(w http.ResponseWriter, req *http.Request) {
	fmt.Println("logging request method2", req.Method)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
}
