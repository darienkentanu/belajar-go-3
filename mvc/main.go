package main

import (
	"fmt"
	"net/http"

	"belajar-go-3/config-file/config"
	"belajar-go-3/mvc/controllers"
	"belajar-go-3/mvc/models"

	"github.com/gorilla/mux"
)

var users []models.User

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/users", controllers.GetUsersEndPoint).Methods("GET")
	router.HandleFunc("/api/users", controllers.CreateUserEndPoint).Methods("POST")
	router.HandleFunc("/api/users/{id}", controllers.GetUserEndPoint).Methods("GET")

	fmt.Println("server started at" + config.GetString("server.address"))
	http.ListenAndServe(config.GetString("server.address"), router)
}
