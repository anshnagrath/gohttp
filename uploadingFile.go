package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", d)
	http.ListenAndServe(":8080", nil)
}
func d(w http.ResponseWriter, req *http.Request) {
	var s string
	if req.Method == http.MethodPost {
		f, h, err := req.FormFile("q")
		if err != nil {
			http.Error(w, "Error While Parsing files", 404)
			return
		}
		defer f.Close()
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
		dst, err := os.Create(filepath.Join("./user/", h.Filename))
		if err != nil {
			fmt.Println(err)
		}
		defer dst.Close()
		numBy, err := dst.Write(bs)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(numBy)

	}
	w.Header().Set("Content-type", "text/html;charset=UTF-8")
	io.WriteString(w, `<form method="POST" enctype="multipart/form-data">
	<input type="File" name="q">
	<input type="Submit">
	</form>
	<br>
	`+s)
}
