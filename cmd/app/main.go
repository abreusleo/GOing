package main

import (
	"GOing/internal/user"
	"GOing/pkg/config"
	"GOing/pkg/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()

	userService := user.NewService(cfg)

	userController := controllers.NewUserController(userService)

	r := mux.NewRouter()
	r.HandleFunc("/users", userController.QueryUsersHandler).Methods("GET")

	log.Println("Servidor iniciado na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
