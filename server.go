package main

import (
	"app/interfaces/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user/create/{name}", userCreateHandler)

	http.ListenAndServe(":8080", r)

	// e.POST("/users", func(c *gin.Context) { userController.Create(c) })
	// e.GET("/users", func(c *gin.Context) { userController.Index(c) })
	// e.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })

	// e = e
}

func userCreateHandler(w http.ResponseWriter, r *http.Request) {
	userController := controllers.NewUserController()
	vars := mux.Vars(r)
	user := userController.Create(vars["name"])
	w.Write([]byte(user.FirstName + user.LastName))
}
