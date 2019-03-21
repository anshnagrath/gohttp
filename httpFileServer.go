package main

import (
	"net/http"
)

func main() {
	// http.ResponseWriter.Header.Set("Content-type", "text/html")
	http.Handle("/test/", http.StripPrefix("/test", http.FileServer(http.Dir("."))))
	http.ListenAndServe(":8080", nil)
}
