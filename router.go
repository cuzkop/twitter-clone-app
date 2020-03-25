package main

import (
	"github.com/kazuki5555/twitter-clone-app/controllers"
	"encoding/json"
	"log"
	"net/http"

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

func timelineHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	timelines, err := controllers.NewTimeline(r)
	if err != nil {
		e := Error{ErrorMsg: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(e)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(timelines)
}

func tweetCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := controllers.NewTweets(r)
	if err != nil {
		e := Error{ErrorMsg: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(e)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func tweetDeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := controllers.NewTweets(r)
	if err != nil {
		e := Error{ErrorMsg: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(e)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func commentCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := controllers.NewComment(r)
	if err != nil {
		e := Error{ErrorMsg: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func favoriteCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := controllers.NewFavorites(r)
	if err != nil {
		e := Error{ErrorMsg: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(e)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func favoriteDeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := controllers.NewFavorites(r)
	if err != nil {
		e := Error{ErrorMsg: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(e)
		return
	}

	w.WriteHeader(http.StatusOK)
}
