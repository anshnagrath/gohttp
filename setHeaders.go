package main

import (
	"fmt"
	"net/http"
)

type test int

func (m test) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("mykey", "myvalue")
	res.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprintln(res, "ascdascasdcascasdc")
}
func main() {
	var d test
	http.ListenAndServe(":8080", d)
}
