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
	m := models.NewSqlHandler()
	users, err := controllers.NewUsers(mux.Vars(r)["screen_id"], m)

	timeline := controllers.NewTimeline()
	timelines, err := timeline.GetTimeline(users, m)
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

	users, err := controllers.NewUsers(mux.Vars(r)["screen_id"], m)
	if err != nil {
		e := Error{ErrorMsg: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
		return
	}

	tweets := controllers.NewTweets()
	tweets.UserID = users.ID

	err = json.NewDecoder(r.Body).Decode(&tweets)
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

	users, err := controllers.NewUsers(mux.Vars(r)["screen_id"], m)
	if err != nil {
		e := Error{ErrorMsg: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
		return
	}

	tweets := controllers.NewTweets()
	tweets.UserID = users.ID

	err = json.NewDecoder(r.Body).Decode(&tweets)
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

func commentCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	m := models.NewSqlHandler()

	users, err := controllers.NewUsers(mux.Vars(r)["screen_id"], m)
	if err != nil {
		e := Error{ErrorMsg: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
		return
	}

	tweets := controllers.NewTweets()
	tweets.UserID = users.ID
	tweets.IsComment = 1
	tweets.TweetID, _ = strconv.Atoi(mux.Vars(r)["tweet_id"])

	err = json.NewDecoder(r.Body).Decode(&tweets)
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

func favoriteCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	m := models.NewSqlHandler()

	users, err := controllers.NewUsers(mux.Vars(r)["screen_id"], m)
	if err != nil {
		e := Error{ErrorMsg: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
		return
	}

	favorites := controllers.NewFavorites()
	favorites.UserID = users.ID

	err = json.NewDecoder(r.Body).Decode(&favorites)
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

	users, err := controllers.NewUsers(mux.Vars(r)["screen_id"], m)
	if err != nil {
		e := Error{ErrorMsg: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
		return
	}

	favorites := controllers.NewFavorites()
	favorites.UserID = users.ID

	err = json.NewDecoder(r.Body).Decode(&favorites)
	if err != nil {
		e := Error{ErrorMsg: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
		return
	}

	log.Println(favorites)

	err = controllers.DeleteFavorite(favorites, m)
	if err != nil {
		e := Error{ErrorMsg: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(e)
		return
	}

	w.WriteHeader(http.StatusOK)
}
