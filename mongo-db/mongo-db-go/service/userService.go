package service

import (
	"mongo-db-go/domain"
	"mongo-db-go/dto"
)

type IUserService interface {
	CreateUser(dto.User) (*dto.User, error)
	FindUser(string) (*dto.User, error)
	DeleteUser(string) (*string, error)
	FindUsers() ([]dto.User, error)
}

type LocalUserService struct {
	repo domain.IUserRepo
}

func (lus LocalUserService) CreateUser(user dto.User) (*dto.User, error) {
	return lus.repo.CreateUser(user)
}

func (lus LocalUserService) FindUser(id string) (*dto.User, error) {
	return lus.repo.FindUser(id)
}

func (lus LocalUserService) DeleteUser(id string) (*string, error) {
	return lus.repo.DeleteUser(id)
}

func (lus LocalUserService) FindUsers() ([]dto.User, error) {
	return lus.repo.FindUsers()
}

func NewLocalUserService(repo domain.IUserRepo) LocalUserService {
	return LocalUserService{
		repo: repo,
	}
}
