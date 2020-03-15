package models

import (
	"log"
)

type Favorites struct {
	ID      int `json:"id"`
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

// func (t Tweets) IsDeleteTweet(m *DB) error {
// 	afterTweet := t
// 	f := m.DB.First(&t)
// 	if f.Error != nil {
// 		log.Println(f.Error)
// 		return f.Error
// 	}

// 	afterTweet.IsDeleted = 1
// 	afterTweet.UpdatedAt = time.Now().UTC()

// 	result := m.DB.Model(&t).Update(afterTweet)
// 	if result.Error != nil {
// 		log.Println(result.Error)
// 		return result.Error
// 	}
// 	return nil
// }
