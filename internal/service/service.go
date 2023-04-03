package service

import (
	"one-lab/internal/model"
)

type UserService struct {
	st UserStorage
}

type UserStorage interface {
	Add(id string, u model.UserEntity) error
}

func NewUserService(st UserStorage) *UserService {
	return &UserService{st: st}
}

func (s *UserService) Add(id string, u model.UserEntity) error {
	return s.Add(id, u)
}
