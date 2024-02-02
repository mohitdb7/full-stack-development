package app

import (
	"email-verify-go/domain"
	"email-verify-go/handlers"
	"email-verify-go/service"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func StartApp() {
	router := mux.NewRouter()

	handler := handlers.NewVerifyEmailHandler(
		service.NewVerifyEmailServiceImpl(
			domain.NewVerifyEmailImpl(),
		))

	router.HandleFunc("/verify_email", handler.VerifyEmail).Methods(http.MethodPost)

	fmt.Println("Starting the server")
	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		panic("Cannot start the server")
	}
}
