package controllers

import (
	"GOing/internal/user"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type UserController struct {
	service *user.Service
}

func NewUserController(service *user.Service) *UserController {
	return &UserController{service: service}
}

func (c *UserController) QueryUsersHandler(w http.ResponseWriter, r *http.Request) {
	ids := r.URL.Query().Get("ids")
	if ids == "" {
		http.Error(w, "Parâmetro 'ids' é obrigatório", http.StatusBadRequest)
		return
	}

	userIDs := parseIDs(ids)

	results := c.service.QueryUsers(userIDs)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func parseIDs(ids string) []int {
	idStrs := strings.Split(ids, ",")
	var userIDs []int
	for _, idStr := range idStrs {
		var id int
		fmt.Sscanf(idStr, "%d", &id)
		userIDs = append(userIDs, id)
	}
	return userIDs
}
