package controllers

import (
	"app/models"
)

type Timeline struct {
	TweetID    int    `json:"tweet_id"`
	UserID     int    `json:"user_id"`
	ScreenName string `json:"screen_name"`
	ScreenID   string `json:"screen_id"`
	Text       string `json:"text"`
	CreatedAt  string `json:"created_at"`
}

func NewTimeline() Timeline {
	return Timeline{}
}

func (t *Timeline) GetTimeline(users models.Users, m *models.DB) ([]Timeline, error) {
	followers, err := models.GetFollowingByUid(users.ID, m)
	if err != nil {
		return nil, err
	}

	//timeline取得のために自分も入れる
	userIds := []int{users.ID}
	for _, s := range followers {
		userIds = append(userIds, s.FollowedID)
	}

	var timelines []Timeline
	rows, err := models.GetTimelineByUserId(userIds, m)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		m.DB.ScanRows(rows, &t)
		timelines = append(timelines, *t)
	}
	return timelines, nil
}
