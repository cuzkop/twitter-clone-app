package main

import (
	"log"
	"net/http"
)

func init() {
	log.SetPrefix("[main]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	log.Printf("main start.")
	r := buildRouter()

	http.ListenAndServe(":8080", r)
}
