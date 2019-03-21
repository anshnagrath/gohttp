package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/image", openFile)
	http.ListenAndServe(":8080", nil)
}

//func d(w http.ResponseWriter, req *http.Request) {
// 	w.Header().Set("Content-type", "text/html;charset=UTF-8")
// 	io.WriteString(w, `<img src='./golang.jpg'>`)
// }
func openFile(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("golang.jpg")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	io.Copy(w, f)

}
