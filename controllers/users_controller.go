package controllers

import (
	"github.com/kazuki5555/twitter-clone-app/models"
	"log"
)

func init() {
	log.SetPrefix("[controllers/users]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func NewUsers(screen_id string, m *models.DB) (models.Users, error) {
	users := models.GetUsers()
	users.ScreenId = screen_id
	user, err := models.GetUsersByScreenID(users.ScreenId, m)
	if err != nil {
		return user, err
	}

	return user, nil
}
