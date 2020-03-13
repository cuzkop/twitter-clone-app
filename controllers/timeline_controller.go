package controllers

import (
	"app/models"
	"log"
)

type Timeline struct {
	TweetId    int    `json:"tweet_id"`
	UserId     int    `json:"user_id"`
	ScreenName string `json:"screen_name"`
	Text       string `json:"text"`
	CreatedAt  string `json:"created_at"`
}

type Timelines []Timeline

func NewTimelineController() Timeline {
	return Timeline{}
}

func (t *Timeline) GetTimeline(userId int, m *models.DB) *Timeline {
	t.UserId = userId
	followers := models.GetFollowingByUid(t.UserId, m)

	//timeline取得のために自分も入れる
	userIds := []int{userId}
	for _, s := range followers {
		userIds = append(userIds, s.FollowedId)
	}
	log.Println(userIds)

	rows := models.GetTimelineByUserId(userIds, m)
	for rows.Next() {
		var id int
		var screenName string
		var tweetId int
		var text string
		var createdAt string
		if err := rows.Scan(&id, &screenName, &tweetId, &text, &createdAt); err != nil {
			continue
		}
		timeline := Timeline{
			TweetId:    tweetId,
			UserId:     id,
			ScreenName: screenName,
			Text:       text,
			CreatedAt:  createdAt,
		}
		log.Println(timeline)
	}
	return t
}
