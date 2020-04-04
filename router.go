package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kazuki5555/twitter-clone-app/controllers"

	"github.com/gorilla/mux"
)

type Error struct {
	ErrorMsg string
}

func init() {
	log.SetPrefix("[router]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func buildRouter() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/timeline/{screen_id}", timelineHandler).Methods("GET")
	r.HandleFunc("/tweets/{screen_id}", tweetCreateHandler).Methods("POST")
	r.HandleFunc("/tweets/{screen_id}", tweetDeleteHandler).Methods("DELETE")
	r.HandleFunc("/tweets/{screen_id}/{tweet_id}", commentCreateHandler).Methods("POST")
	r.HandleFunc("/favorites/{screen_id}", favoriteCreateHandler).Methods("POST")
	r.HandleFunc("/favorites/{screen_id}", favoriteDeleteHandler).Methods("DELETE")
	return r
}

func createResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func timelineHandler(w http.ResponseWriter, r *http.Request) {
	timelines, err := controllers.NewTimeline(r)
	if err != nil {
		e := Error{ErrorMsg: err.Error()}
		createResponse(w, http.StatusInternalServerError, e)
		return
	}

	createResponse(w, http.StatusOK, timelines)
}

func tweetCreateHandler(w http.ResponseWriter, r *http.Request) {
	err := controllers.CreateTweet(r)
	if err != nil {
		e := Error{ErrorMsg: err.Error()}
		createResponse(w, http.StatusInternalServerError, e)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func tweetDeleteHandler(w http.ResponseWriter, r *http.Request) {
	err := controllers.IsDeleteTweet(r)
	if err != nil {
		e := Error{ErrorMsg: err.Error()}
		createResponse(w, http.StatusInternalServerError, e)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func commentCreateHandler(w http.ResponseWriter, r *http.Request) {
	err := controllers.NewComment(r)
	if err != nil {
		e := Error{ErrorMsg: err.Error()}
		createResponse(w, http.StatusInternalServerError, e)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func favoriteCreateHandler(w http.ResponseWriter, r *http.Request) {
	err := controllers.CreateFavorite(r)
	if err != nil {
		e := Error{ErrorMsg: err.Error()}
		createResponse(w, http.StatusInternalServerError, e)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func favoriteDeleteHandler(w http.ResponseWriter, r *http.Request) {
	err := controllers.DeleteFavorite(r)
	if err != nil {
		e := Error{ErrorMsg: err.Error()}
		createResponse(w, http.StatusInternalServerError, e)
		return
	}

	w.WriteHeader(http.StatusOK)
}
