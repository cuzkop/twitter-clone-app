package models

import (
	"log"
)

type Favorites struct {
	UserID  int `json:"user_id"`
	TweetId int `json:"tweet_id"`
}

func init() {
	log.SetPrefix("[tweets]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func GetFavorites() Favorites {
	return Favorites{}
}

func (f Favorites) CreateFavorite(m *DB) error {
	result := m.DB.Create(&f)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func (f Favorites) DeleteTweet(m *DB) error {
	result := m.DB.Delete(&f)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}
