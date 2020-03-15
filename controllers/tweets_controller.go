package controllers

import (
	"app/models"
)

func NewTweets() models.Tweets {
	return models.GetTweets()
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
