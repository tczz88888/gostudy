package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandle(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Url.Path = %q\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/hello", helloHandle)
	http.HandleFunc("/", indexHandle)
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}
