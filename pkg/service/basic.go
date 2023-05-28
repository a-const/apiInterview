package service

import (
	"apiInterview/pkg/repo"
	user "apiInterview/pkg/type"
)

type Basic interface {
	CreateUser(username string, password string, decription string) error
	GetUser(username string) (*user.User, error)
	UpdateUser(username string, password string, decription string) error
	DeleteUser(username string) error
	GetAllUsers() (*[]user.User, error)
}

type Service struct {
	Basic
}

func NewService(repo *repo.MongoDB) *Service {
	return &Service{
		Basic: NewBasicService(repo),
	}
}
