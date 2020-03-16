package models

import (
	"log"
	"time"
)

type Tweets struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Text      string    `json:"text"`
	TweetID   int       `json:"tweet_id"`
	IsComment int       `json:"is_comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted int       `json:"is_deleted"`
}

func init() {
	log.SetPrefix("[tweets]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func GetTweets() Tweets {
	return Tweets{}
}

func (t Tweets) CreateTweet(m *DB) error {
	result := m.DB.Create(&t)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func (t Tweets) IsDeleteTweet(m *DB) error {
	afterTweet := t
	f := m.DB.First(&t)
	if f.Error != nil {
		log.Println(f.Error)
		return f.Error
	}

	afterTweet.IsDeleted = 1
	afterTweet.UpdatedAt = time.Now().UTC()

	result := m.DB.Model(&t).Update(afterTweet)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}
