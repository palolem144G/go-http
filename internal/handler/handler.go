package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type UserStore map[int]User

var userStore = make(UserStore)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newUser User

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newId := len(userStore) + 1
	newUser.Id = newId

	userStore[newId] = newUser

	w.WriteHeader(http.StatusCreated)
	message := fmt.Sprintf("User id:%d, name:%s, role: %s, created", newUser.Id, newUser.Name, newUser.Role)
	w.Write([]byte(message))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	user, ok := userStore[id]
	log.Println(user)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		errMsg := fmt.Sprintf("User with id %d does not exist", id)
		w.Write([]byte(errMsg))
		return
	}

	respBody, err := json.Marshal(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	uiid := r.URL.Query().Get("id")
	id, err := strconv.Atoi(uiid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	delete(userStore, id)
	if err != nil {
		errMsg := fmt.Sprintf("User with id %d has not been deleted", id)
		w.Write([]byte(errMsg))
		return
	}
	w.WriteHeader(http.StatusOK)

}

// Написать методы Delete (удалить юзера по id) и GetAll (возвращает всех юзеров)
