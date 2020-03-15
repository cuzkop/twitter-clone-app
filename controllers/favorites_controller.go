package controllers

import (
	"app/models"
)

func NewFavorites() models.Favorites {
	return models.GetFavorites()
}

func CreateFavorite(f models.Favorites, m *models.DB) error {
	err := f.CreateFavorite(m)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFavorite(f models.Favorites, m *models.DB) error {
	err := f.DeleteTweet(m)
	if err != nil {
		return err
	}
	return nil
}
