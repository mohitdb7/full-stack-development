package service

import (
	"mongo-db-go/domain"
	"mongo-db-go/dto"
)

type IUserService interface {
	CreateUser(dto.User) (*dto.User, error)
}

type LocalUserService struct {
	repo domain.IUserRepo
}

func (lus LocalUserService) CreateUser(user dto.User) (*dto.User, error) {
	return lus.repo.CreateUser(user)
}

func NewLocalUserService(repo domain.IUserRepo) LocalUserService {
	return LocalUserService{
		repo: repo,
	}
}
