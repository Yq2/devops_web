package main

import (
	"net/http"
	"io"
)

func firstPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>杨强的一个DevOps实例</h1>")
}

func main () {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":8000", nil)
}
