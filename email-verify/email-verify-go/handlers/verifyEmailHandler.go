package handlers

import (
	"email-verify-go/dto"
	"email-verify-go/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type VerifyEmailHandler struct {
	service service.IVerifyEmailService
}

func (veh VerifyEmailHandler) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	emailModel := dto.EmailModel{}
	err := json.NewDecoder(r.Body).Decode(&emailModel)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unable to read the object %v", err)
		return
	}

	emailVerify := veh.service.IsEmailValid(emailModel.EmailDomain)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(emailVerify); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unable to write object %v", err)
	}
}

func NewVerifyEmailHandler(service service.IVerifyEmailService) VerifyEmailHandler {
	return VerifyEmailHandler{
		service: service,
	}
}
