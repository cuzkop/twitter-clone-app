package controllers

import (
	"github.com/kazuki5555/twitter-clone-app/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func init() {
	log.SetPrefix("[controllers/tweets]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func NewTweets(r *http.Request) error {
	m := models.NewSqlHandler()
	users, err := NewUsers(mux.Vars(r)["screen_id"], m)
	if err != nil {
		return err
	}

	tweets := models.GetTweets()
	tweets.UserID = users.ID

	err = json.NewDecoder(r.Body).Decode(&tweets)
	if err != nil {
		log.Println(err)
		return err
	}

	if r.Method == "POST" {
		err = CreateTweet(tweets, m)
	} else if r.Method == "DELETE" {
		err = IsDeleteTweet(tweets, m)
	}
	if err != nil {
		return err
	}

	return nil
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

	err = CreateTweet(tweets, m)
	if err != nil {
		return err
	}

	return nil
}

func CreateTweet(t models.Tweets, m *models.DB) error {
	err := t.CreateTweet(m)
	if err != nil {
		return err
	}
	return nil
}

func IsDeleteTweet(t models.Tweets, m *models.DB) error {
	err := t.IsDeleteTweet(m)
	if err != nil {
		return err
	}
	return nil
}
