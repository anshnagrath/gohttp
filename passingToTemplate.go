package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type formValues struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))

}
func main() {
	http.HandleFunc("/", test)
	http.ListenAndServe(":8080", nil)
}
func test(w http.ResponseWriter, req *http.Request) {

	f := req.FormValue("first")
	l := req.FormValue("last")
	s := req.FormValue("subscribed") == "on"
	p := formValues{
		f,
		l,
		s,
	}
	tpl.ExecuteTemplate(w, "passingToTemplate.html", p)
	fmt.Println("%v", f, l, s)
}
