package app

import (
	"file-upload-go/domain"
	"file-upload-go/handlers"
	"file-upload-go/services"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func dummyFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Dummy function")
}

func StartApp() {
	// 1. Create the router
	router := mux.NewRouter()

	handler := handlers.NewUploadHandler(
		services.NewLocalUploadService(
			domain.NewLocalUpload(),
		))

	// 2. Add functions to router
	router.HandleFunc("/upload", handler.FileUpload).Methods(http.MethodPost)
	router.HandleFunc("/dummy", dummyFunc)

	// 3. Start the server
	fmt.Println("Server started")
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		panic(fmt.Sprintf("Unable to start the server: %s", err))
	}
}
