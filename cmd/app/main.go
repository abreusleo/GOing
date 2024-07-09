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
	// Carregar configurações do ambiente
	cfg := config.LoadConfig()

	// Inicializar o serviço de usuários
	userService := user.NewService(cfg)

	userController := controllers.NewUserController(userService)

	// Configurar o roteador
	r := mux.NewRouter()
	r.HandleFunc("/users", userController.QueryUsersHandler).Methods("GET")

	// Iniciar o servidor HTTP
	log.Println("Servidor iniciado na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
