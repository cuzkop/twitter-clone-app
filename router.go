package main

import (
	"app/controllers"
	"app/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Error struct {
	ErrorMsg string
}

func init() {
	log.SetPrefix("[router]]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func buildRouter() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/timeline/{user_id}", timelineHandler).Methods("GET")
	r.HandleFunc("/tweets/{user_id}", tweetCreateHandler).Methods("POST")
	r.HandleFunc("/tweets/{user_id}", tweetDeleteHandler).Methods("DELETE")
	r.HandleFunc("/favorites/{user_id}", favoriteCreateHandler).Methods("POST")
	r.HandleFunc("/favorites/{user_id}", favoriteDeleteHandler).Methods("DELETE")
	// r.HandleFunc("/comment/{name}", userCreateHandler)
	return r
}

func timelineHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	m := models.NewSqlHandler()
	timeline := controllers.NewTimeline()
	user_id, _ := strconv.Atoi(mux.Vars(r)["user_id"])
	timelines, err := timeline.GetTimeline(user_id, m)
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
	m := models.NewSqlHandler()
	tweets := controllers.NewTweets()
	tweets.UserID, _ = strconv.Atoi(mux.Vars(r)["user_id"])

	err := json.NewDecoder(r.Body).Decode(&tweets)
	if err != nil {
		e := Error{ErrorMsg: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
		return
	}

	err = controllers.CreateTweet(tweets, m)
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
	m := models.NewSqlHandler()
	tweets := controllers.NewTweets()
	tweets.UserID, _ = strconv.Atoi(mux.Vars(r)["user_id"])

	err := json.NewDecoder(r.Body).Decode(&tweets)
	if err != nil {
		e := Error{ErrorMsg: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
		return
	}

	err = controllers.IsDeleteTweet(tweets, m)
	if err != nil {
		e := Error{ErrorMsg: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(e)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func favoriteCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	m := models.NewSqlHandler()
	favorites := controllers.NewFavorites()
	favorites.UserID, _ = strconv.Atoi(mux.Vars(r)["user_id"])

	err := json.NewDecoder(r.Body).Decode(&favorites)
	if err != nil {
		e := Error{ErrorMsg: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
		return
	}

	err = controllers.CreateFavorite(favorites, m)
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
	m := models.NewSqlHandler()
	favorites := controllers.NewFavorites()
	favorites.UserID, _ = strconv.Atoi(mux.Vars(r)["user_id"])

	err := json.NewDecoder(r.Body).Decode(&favorites)
	if err != nil {
		e := Error{ErrorMsg: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
		return
	}

	err = controllers.DeleteFavorite(favorites, m)
	if err != nil {
		e := Error{ErrorMsg: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(e)
		return
	}

	w.WriteHeader(http.StatusOK)
}
