package service

import (
	"email-verify-go/domain"
	"email-verify-go/dto"
)

type IVerifyEmailService interface {
	IsEmailValid(string) dto.EmailVerifyReponseModel
}

type VerifyEmailServiceImpl struct {
	repo domain.IVerifyEmail
}

func (vesi VerifyEmailServiceImpl) IsEmailValid(email string) dto.EmailVerifyReponseModel {
	return vesi.repo.IsEmailValid(email)
}

func NewVerifyEmailServiceImpl(repo domain.IVerifyEmail) IVerifyEmailService {
	return VerifyEmailServiceImpl{
		repo: repo,
	}
}
