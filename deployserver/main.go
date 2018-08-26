package main

import (
	"net/http"
	"io"
	"os/exec"
	"log"
)

func reLaunch() {
	cmd := exec.Command("sh","./deploy.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func firstPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h2> webserver DevOps 入口</h2>")
	reLaunch()
}

func main () {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":5000", nil)
}
