package main

import (
	"io"
	"log"
	"net/http"
)
func serveIndex (w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}


func main() {
	http.HandleFunc("/", serveIndex)

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}