package handlers

import (
	"file-upload-go/services"
	"fmt"
	"net/http"
)

type UploadHandler struct {
	service services.IUploadService
}

func (uh UploadHandler) FileUpload(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Reached upload\n")

	//1. Param input for multipart file upload
	r.ParseMultipartForm(200 << 20) // Maximum of 200MB file allowed

	//2. Retrieve file from form-data
	//<Form-id> is the form key that we will read from. Client should use the same form key when uploading the file
	file, handler, err := r.FormFile("form-id")
	if err != nil {
		errStr := fmt.Sprintf("Error in reading the file %s\n", err)
		fmt.Println(errStr)
		fmt.Fprintf(w, errStr)
		return
	}

	result, err := uh.service.SaveFile(file, handler)
	fmt.Fprintf(w, result)

	if err != nil {
		// Error handling here
		return
	}
}

func NewUploadHandler(service services.IUploadService) UploadHandler {
	return UploadHandler{
		service: service,
	}
}
