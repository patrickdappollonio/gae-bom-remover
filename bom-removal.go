package main

import (
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/upload", upload)
	log.Println("Listening on port 8080")
}
