package main

import (
	"net/http"
	"io"
)

func firstPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h2>devops_web </h2>")
}

func main () {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":8000", nil)
}

// 1. git pull
//2 git push -> git pull
//3.deploy
