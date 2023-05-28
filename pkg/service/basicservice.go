package service

import (
	"apiInterview/pkg/repo"
	user "apiInterview/pkg/type"
	"crypto/sha1"
	"fmt"
)

type BasicService struct {
	repo *repo.MongoDB
}

const salt = "dfvdfvjiurfnu934f3nlakmsi"

func NewBasicService(repo *repo.MongoDB) *BasicService {
	return &BasicService{repo: repo}
}

func (bs *BasicService) CreateUser(username string, password string, decription string) error {
	passHash := generatePassHash(password)
	newUser := &user.User{
		Username:    username,
		Password:    passHash,
		Description: decription,
	}
	return bs.repo.CreateUser(newUser)
}

func (bs *BasicService) GetUser(username string) (*user.User, error) {
	return bs.repo.GetUser(username)
}

func (bs *BasicService) UpdateUser(username string, password string, decription string) error {
	return bs.repo.UpdateUser(username, generatePassHash(password), decription)
}

func (bs *BasicService) DeleteUser(username string) error {
	return bs.repo.DeleteUser(username)
}

func (bs *BasicService) GetAllUsers() (*[]user.User, error) {
	return bs.repo.GetAllUsers()
}

func generatePassHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
