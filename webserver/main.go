package main

import (
	"net/http"
	"io"
)

func firstPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h2> Yq2 devops_web v1.0.9 </h2>")
}

func main () {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":8000", nil)
}

// 1. git pull
//2 git push -> git pull
//3.deploy
