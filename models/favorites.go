package models

import (
	"log"
)

type Favorites struct {
	UserID  int `json:"user_id"`
	TweetId int `json:"tweet_id"`
}

func init() {
	log.SetPrefix("[models/favorites]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func GetFavorites() Favorites {
	return Favorites{}
}

func (f Favorites) CreateFavorite(m *DB) error {
    tx := m.DB.Begin()
	result := tx.Create(&f)
	if result.Error != nil {
        tx.Rollback()
		log.Println(result.Error)
		return result.Error
    }
    tx.Commit()
	return nil
}

func (f Favorites) DeleteFavorite(m *DB) error {
	result := m.DB.Where("user_id = ? and tweet_id = ?", f.UserID, f.TweetId).Delete(&f)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}
