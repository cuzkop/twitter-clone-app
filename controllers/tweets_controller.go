package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/kazuki5555/twitter-clone-app/models"

	"github.com/gorilla/mux"
)

func init() {
	log.SetPrefix("[controllers/tweets]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func NewComment(r *http.Request) error {
	m := models.NewSqlHandler()
	users, err := NewUsers(mux.Vars(r)["screen_id"], m)
	if err != nil {
		return err
	}

	tweets := models.GetTweets()
	tweets.UserID = users.ID
	tweets.IsComment = 1
	tweets.TweetID, _ = strconv.Atoi(mux.Vars(r)["tweet_id"])

	err = json.NewDecoder(r.Body).Decode(&tweets)
	if err != nil {
		log.Println(err)
		return err
	}

	return tweets.CreateTweet(m)
}

func CreateTweet(r *http.Request) error {
	t, m, err := PrepareTweet(r)
	if err != nil {
		return err
	}

	return t.CreateTweet(m)
}

func IsDeleteTweet(r *http.Request) error {
	t, m, err := PrepareTweet(r)
	if err != nil {
		return err
	}

	return t.IsDeleteTweet(m)
}

func PrepareTweet(r *http.Request) (models.Tweets, *models.DB, error) {
	m := models.NewSqlHandler()
	tweets := models.GetTweets()
	users, err := NewUsers(mux.Vars(r)["screen_id"], m)
	if err != nil {
		return tweets, m, err
	}

	tweets.UserID = users.ID

	err = json.NewDecoder(r.Body).Decode(&tweets)
	if err != nil {
		log.Println(err)
		return tweets, m, err
	}
	return tweets, m, nil
}
