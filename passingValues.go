package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", test)
	http.ListenAndServe(":8080", nil)
}
func test(w http.ResponseWriter, req *http.Request) {
	fname := req.FormValue("fname")
	lname := req.FormValue("lname")
	w.Header().Set("Content-type", "text/html;charset=UTF-8")
	io.WriteString(w, `<!DOCTYPE html>
	<html>
	<head>
		<meta charset="utf-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<title>Http test</title>
		<meta name="viewport" content="width=device-width, initial-scale=1">    
	</head>
	<body>
	<div>
	   <div style="margin-top:5%">
	   <form method="POST">
		   <input type="text"name="fname" placeholder="first name">
		   <input type="text"name="lname" placeholder="last name">
		   <input type="submit" placeholder="Submit">
		</form>   
		</div>
	</body>
	</html>`+fname+" "+lname)
}
