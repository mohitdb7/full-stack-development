package handlers

import (
	"encoding/json"
	"fmt"
	"mongo-db-go/dto"
	"mongo-db-go/service"
	"net/http"
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

func NewUserHandler(service service.IUserService) UserHandler {
	return UserHandler{
		service: service,
	}
}
