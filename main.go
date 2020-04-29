package main

import (
	"log"
	"net/http"

	"github.com/kazuki5555/twitter-clone-app/models"
)

func init() {
	log.SetPrefix("[main]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	models.NewSqlHandler()

}

func main() {
	log.Printf("main start.")
	log.Fatal(http.ListenAndServe(":8080", buildRouter()))
}
