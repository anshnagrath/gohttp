package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/expire", exp)
	http.ListenAndServe(":8080", nil)
}
func set(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("counter")
	fmt.Println("%v", err, 1)
	if err != nil {
		http.SetCookie(w, &http.Cookie{
			Name:  "counter",
			Value: "1",
		})
		fmt.Fprint(w, `<p> Your Cookie value : </p>`+"1")
	}
	fmt.Println("%v", err, 2)
	if err == nil {
		fmt.Println("fcasdcas")
		count, err := strconv.Atoi(c.Value)
		if err != nil {
			log.Fatalln(err)
		}
		count += 1
		c.Value = strconv.Itoa(count)
		http.SetCookie(w, &http.Cookie{
			Name:  "counter",
			Value: c.Value,
		})
		fmt.Fprint(w, `<p> Your Cookie value : </p>`+c.Value)
	}

}
func exp(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("counter")
	fmt.Println("%v", err, 1)
	if err != nil {
		fmt.Fprint(w, `<p> Cookie not Found </p>`)
	}
	c.MaxAge = -1
	http.SetCookie(w, c)
	fmt.Fprint(w, `<p> Cookies Removed </p>`)

}
