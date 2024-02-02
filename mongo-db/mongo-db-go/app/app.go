package app

import (
	"fmt"
	"mongo-db-go/database"
	"mongo-db-go/domain"
	"mongo-db-go/handlers"
	"mongo-db-go/service"
	"net/http"

	"github.com/gorilla/mux"
)

func Startup() {
	router := mux.NewRouter()

	collection, client, context := database.SetupMongoDB()

	handler := handlers.NewUserHandler(
		service.NewLocalUserService(
			domain.NewUserRepoMongo(collection, client, context)))

	router.HandleFunc("/createUser", handler.CreateUser).Methods(http.MethodPost)

	fmt.Println("Running at localhost:8080")
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		panic("Cannot start the server")
	}
}
