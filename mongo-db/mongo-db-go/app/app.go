package app

import (
	"fmt"
	"mongo-db-go/domain"
	"mongo-db-go/handlers"
	"mongo-db-go/service"
	"net/http"

	"github.com/gorilla/mux"
)

func Startup() {
	router := mux.NewRouter()

	handler := handlers.NewUserHandler(
		service.NewLocalUserService(
			domain.NewUserRepoMongo()))

	router.HandleFunc("/create_user", handler.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{user_id}", handler.FindUser).Methods(http.MethodGet)
	router.HandleFunc("/users", handler.FindUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/delete/{user_id}", handler.DeleteUser).Methods(http.MethodGet)

	fmt.Println("Running at localhost:8080")
	//192.168.1.3
	//172.16.32.1
	err := http.ListenAndServe("192.168.1.3:8080", router)
	if err != nil {
		panic("Cannot start the server")
	}
}
