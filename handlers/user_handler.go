package handlers

import (
	"encoding/json"
	"net/http"
	"taskmanager/models"
)

var users = make(map[int]models.User)
var userIDCounter = 1


func GetUsers(w http.ResponseWriter, r *http.Request) {
	userList := make([]models.User, 0, len(users))
	for _, user := range users {
		userList = append(userList, user)
	}
	json.NewEncoder(w).Encode(userList)
}


func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	user.ID = userIDCounter
	userIDCounter++
	users[user.ID] = user
	json.NewEncoder(w).Encode(user)
}
