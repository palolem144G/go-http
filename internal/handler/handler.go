package handler

import (
	"encoding/json"
	"fmt"
	"http-demo/internal/domain"
	"net/http"
	"strconv"
)

type UserService interface {
	CreateUser(user domain.User) (domain.User, error)
	ChangePassword(user domain.User) error
	GetUser(userId int) (domain.User, error)
	DeleteUser(userId int) (domain.User, error)
}

type UserHandler struct {
	us UserService
}

func NewUserHandler(us UserService) UserHandler {
	return UserHandler{
		us: us,
	}
}

func (uh UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newUser domain.User
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createdUser, err := uh.us.CreateUser(newUser)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "test/plain")
	w.WriteHeader(http.StatusCreated)
	message := fmt.Sprintf("User id: %d, name: %s, role: %s, created", createdUser.Id, createdUser.Name, createdUser.Role)
	w.Write([]byte(message))
}

func (uh UserHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user domain.User
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = uh.us.ChangePassword(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (uh UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	user, err := uh.us.GetUser(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		errMsg := fmt.Sprintf("User with id %d does not exist", id)
		w.Write([]byte(errMsg))
		return
	}

	respBody, err := json.Marshal(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

// Method GetAll
// func(un UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		return
// 	}
// }

func (uh UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	uiid := r.URL.Query().Get("id")
	id, err := strconv.Atoi(uiid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	uh.us.DeleteUser(id)
	if err != nil {
		errMsg := fmt.Sprintf("User with id %d has not been deleted", id)
		w.Write([]byte(errMsg))
		return
	}
	w.WriteHeader(http.StatusOK)

}

// Написать методы Delete (удалить юзера по id) и GetAll (возвращает всех юзеров)
