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

func CreateFavorite(r *http.Request) error {
	m := models.NewSqlHandler()

	users, err := NewUsers(mux.Vars(r)["screen_id"], m)
	if err != nil {
		return err
	}

	favorites := models.GetFavorites()
	favorites.UserID = users.ID

	err = json.NewDecoder(r.Body).Decode(&favorites)
	if err != nil {
		log.Println(err)
		return err
	}

	return favorites.CreateFavorite(m)
}

func DeleteFavorite(r *http.Request) error {
	m := models.NewSqlHandler()

	users, err := NewUsers(mux.Vars(r)["screen_id"], m)
	if err != nil {
		return err
	}

	favorites := models.GetFavorites()
	favorites.UserID = users.ID

	err = json.NewDecoder(r.Body).Decode(&favorites)
	if err != nil {
		log.Println(err)
		return err
	}

	return favorites.DeleteFavorite(m)
}
