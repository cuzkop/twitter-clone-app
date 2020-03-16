package models

import (
	"database/sql"
	"log"
)

type Users struct {
	ID         int    `json:"id"`
	ScreenName string `json:"screen_name"`
	ScreenId   string `json:"screen_id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

func init() {
	log.SetPrefix("[models/users]]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func GetUsers() Users {
	return Users{}
}

func GetUsersByScreenID(screen_id string, m *DB) (Users, error) {
	var users Users
	result := m.DB.First(&users, "screen_id = ?", screen_id)
	if result.Error != nil {
		log.Println(result.Error)
		return users, result.Error
	}

	return users, result.Error
}

func GetTimelineByUserId(users []int, m *DB) (*sql.Rows, error) {
	query := m.DB.Table("users").
		Select("users.id as user_id, users.screen_name, users.screen_id, tweets.id as tweet_id, tweets.text, tweets.created_at").
		Joins("left join tweets on users.id = tweets.user_id").
		Where("users.id in (?) and tweets.is_deleted = 0 and tweets.is_comment = 0", users).
		Order("tweets.created_at desc")
	rows, err := query.Rows()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return rows, nil
}
