package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/serveFile", serve)
	http.ListenAndServe(":8080", nil)
}
func serve(w http.ResponseWriter, req *http.Request) {

	http.ServeFile(w, req, "golang.jpg")
}
