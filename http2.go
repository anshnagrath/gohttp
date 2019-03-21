package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

type handler int

var tpl *template.Template

func (m handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// res.Header().Add("Content-Type", "html")
	err := req.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	data := struct {
		Method        string
		Url           *url.URL
		Header        http.Header
		Submissions   map[string][]string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Header,
		req.Form,
		req.ContentLength,
	}
	tpl.ExecuteTemplate(res, "http2.html", data)
}
func init() {
	tpl = template.Must(template.ParseFiles("http2.html"))

}
func main() {
	var m handler
	http.ListenAndServe(":8080", m)
}
