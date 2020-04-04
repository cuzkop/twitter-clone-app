package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kazuki5555/twitter-clone-app/models"

	"github.com/gorilla/mux"
)

func init() {
	log.SetPrefix("[controllers/favorites]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func PrepareFavorite(r *http.Request) (models.Favorites, *models.DB, error) {
	m := models.NewSqlHandler()
	favorites := models.GetFavorites()

	users, err := NewUsers(mux.Vars(r)["screen_id"], m)
	if err != nil {
		return favorites, m, err
	}

	favorites.UserID = users.ID

	err = json.NewDecoder(r.Body).Decode(&favorites)
	if err != nil {
		log.Println(err)
		return favorites, m, err
	}
	return favorites, m, nil
}

func CreateFavorite(r *http.Request) error {
	f, m, err := PrepareFavorite(r)
	if err != nil {
		log.Println(err)
		return err
	}

	return f.CreateFavorite(m)
}

func DeleteFavorite(r *http.Request) error {
	f, m, err := PrepareFavorite(r)
	if err != nil {
		log.Println(err)
		return err
	}

	return f.DeleteFavorite(m)
}
