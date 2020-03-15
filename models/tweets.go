package models

import (
	"log"
	"time"
)

const location = "Asia/Tokyo"

type Tweets struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Text      string    `json:"text"`
	CreatedAt string    `json:"created_at"`
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
