package controllers

import (
	"encoding/json"
	"fmt"
	"go/pkg/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"main.go/pkg/utils"

	"main.go/pkg/models"
)

var User models.User

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetUsers()
	res, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	users, _ := models.GetUserById(ID)
	res, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	utils.ParseBody(r, user)
	u := user.CreateUser()
	res, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	users := models.DeleteUser(ID)
	res, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user = &models.User{}
	utils.ParseBody(r, user)
	vars := mux.Vars(r)

	userId := vars["id"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	users, db := models.GetUserById(ID)
	if users.Username != "" {
		users.Username = user.Username
	}
	if users.Email != "" {
		users.Email = user.Email
	}
	if users.Password != "" {
		users.Password = user.Password
	}
	db.Save(&users)
	res, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
