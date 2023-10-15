package main

import (
	"fmt"
	"net/http"
) 
func main() {
	fmt.Println("this project is a simple server initilisation  by golang")
	// this is simple server initilisation by golang 

	html := 		`<html>
						<head>
							<title>Page Title</title>
						</head>
						<body>
							<h1>Welcome to my first server</h1>
						</body>
					</html>`
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, html)
	})


	
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}