package models

import (
	"log"
)

type Followers struct {
	FollowingId int `json:"following_id"`
	FollowedId  int `json:"followed_id"`
}

type FollowersList []Followers

func init() {
	log.SetPrefix("[followers]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func GetFollowingByUid(uid int, m *DB) FollowersList {
	var followersList FollowersList
	m.DB.Find(&followersList, "following_id = ?", uid)

	return followersList
}
