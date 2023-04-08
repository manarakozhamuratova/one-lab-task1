package service

import (
	"one-lab/internal/model"
)

type UserService struct {
	st UserStorage
}

// что бы удобнее определять потом где твой интерфейс лучше называть их ISomeInterface 
type UserStorage interface {
	Add(id string, u model.UserEntity) error
}

func NewUserService(st UserStorage) *UserService {
	return &UserService{st: st}
}

func (s *UserService) Add(id string, u model.UserEntity) error {
	return s.st.Add(id, u)
}
