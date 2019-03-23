package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.icon", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	fmt.Printf("%+v", v)
	io.WriteString(w, "Search item:"+v)
}
