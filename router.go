package main

import (
	"app/controllers"
	"app/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func init() {
	log.SetPrefix("[router]]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func buildRouter() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/timeline/{user_id}", timelineHandler).Methods("GET")
	r.HandleFunc("/tweets/{user_id}", userCreateHandler).Methods("POST")
	r.HandleFunc("/tweets/{user_id}", userCreateHandler).Methods("DELETE")
	r.HandleFunc("/favorite/{user_id}", userCreateHandler).Methods("POST")
	r.HandleFunc("/favorite/{user_id}", userCreateHandler).Methods("DELETE")
	r.HandleFunc("/comment/{name}", userCreateHandler)
	return r
}

func timelineHandler(w http.ResponseWriter, r *http.Request) {
	m := models.NewSqlHandler()
	timelineController := controllers.NewTimelineController()
	vars := mux.Vars(r)
	user_id, _ := strconv.Atoi(vars["user_id"])
	timelines := timelineController.GetTimeline(user_id, m)
	fmt.Println(timelines)
}

func userCreateHandler(w http.ResponseWriter, r *http.Request) {
	userController := controllers.NewUserController()
	vars := mux.Vars(r)
	user := userController.Create(vars["name"])
	// w.Write([]byte(user.FirstName + user.LastName))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(user)
}
