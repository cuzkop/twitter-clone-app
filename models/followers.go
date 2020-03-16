package models

import (
	"log"
)

type Followers struct {
	FollowingID int `json:"following_id"`
	FollowedID  int `json:"followed_id"`
}

type FollowersList []Followers

func init() {
	log.SetPrefix("[followers]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func GetFollowingByUid(uid int, m *DB) (FollowersList, error) {
	var followersList FollowersList
	result := m.DB.Find(&followersList, "following_id = ?", uid)
	if result.Error != nil {
		log.Println(result.Error)
		return followersList, result.Error
	}

	return followersList, result.Error
}
