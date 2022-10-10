package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/headers", headers)

	fmt.Println("Listening on http://127.0.0.1:3000")

	http.ListenAndServe(":3000", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "<h1>Hello World!</h1>")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
