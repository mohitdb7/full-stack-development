package handlers

import (
	"encoding/json"
	"fmt"
	"mongo-db-go/dto"
	"mongo-db-go/service"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	service service.IUserService
}

func (uh UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	u := dto.User{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		fmt.Fprintf(w, "Error in reading post body %s", err)
		return
	}

	user, err := uh.service.CreateUser(u)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("Error in Creating user %s", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(*user); err != nil {
		panic(err)
	}
}

func (uh UserHandler) FindUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["user_id"]

	user, err := uh.service.FindUser(userId)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("Error in Finding user %v", err))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(*user); err != nil {
		panic(err)
	}
}

func (uh UserHandler) FindUsers(w http.ResponseWriter, r *http.Request) {

	users, err := uh.service.FindUsers()
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("Error in Finding user %v", err))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}
}

func (uh UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["user_id"]

	resultString, err := uh.service.DeleteUser(userId)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("Error in Deleting user %v", err))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, *resultString)
}

func NewUserHandler(service service.IUserService) UserHandler {
	return UserHandler{
		service: service,
	}
}
