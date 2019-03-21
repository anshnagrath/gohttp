package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/serveContent", serve)
	http.ListenAndServe(":8080", nil)
}
func serve(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("golang.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	http.ServeContent(w, req, fi.Name(), fi.ModTime(), f)
}
