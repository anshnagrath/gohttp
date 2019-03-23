package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", set)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/expire", exp)
	http.ListenAndServe(":8080", nil)
}
func set(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("	uid")
	fmt.Println("%v", err, 1)
	if err != nil {
		uid, _ := uuid.NewV4()

		http.SetCookie(w, &http.Cookie{
			Name:  uid.String(),
			Value: "1",

			//secure : true  if we wont to sendthis over https only
			HttpOnly: true, //make the cookie acceseable through http only not by js
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
		uid, _ := uuid.NewV4()
		http.SetCookie(w, &http.Cookie{
			Name:  uid.String(),
			Value: c.Value,
			//secure : true  if we wont to sendthis over https only
			HttpOnly: true, //make the cookie acceseable through http only not by js
		})
		fmt.Fprint(w, `<p> Your Cookie value : </p>`+c.Value)
	}

}
func exp(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("uid")
	fmt.Println("%v", err, 1)
	if err != nil {
		fmt.Fprint(w, `<p> Cookie not Found </p>`)
	}
	c.MaxAge = -1
	http.SetCookie(w, c)
	fmt.Fprint(w, `<p> Cookies Removed </p>`)

}
