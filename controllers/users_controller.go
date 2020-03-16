package controllers

import "app/models"

func NewUsers(screen_id string, m *models.DB) (models.Users, error) {
	users := models.GetUsers()
	users.ScreenId = screen_id
	user, err := models.GetUsersByScreenID(users.ScreenId, m)
	if err != nil {
		return user, err
	}

	return user, nil
}
