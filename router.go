package main

import (
	"app/controllers"
	"app/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const dbURL = "root@tcp(127.0.0.1)/twitter"

func init() {
	log.SetPrefix("[router]]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func buildRouter() http.Handler {
	m := models.NewSqlHandler()
	fmt.Println(m.DB)
	r := mux.NewRouter()

	r.HandleFunc("/user/{name}", userCreateHandler)
	return r
}

func userCreateHandler(w http.ResponseWriter, r *http.Request) {
	userController := controllers.NewUserController()
	vars := mux.Vars(r)
	user := userController.Create(vars["name"])
	w.Write([]byte(user.FirstName + user.LastName))
}
