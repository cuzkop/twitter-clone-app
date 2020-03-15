package main

import (
	"log"
	"net/http"
	"time"
)

const location = "Asia/Tokyo"

func init() {
	log.SetPrefix("[main]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

func main() {
	log.Printf("main start.")
	r := buildRouter()

	http.ListenAndServe(":8080", r)
}
