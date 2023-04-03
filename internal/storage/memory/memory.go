package memory

import (
	"one-lab/internal/model"
	"sync"
)

type Storage struct {
	users map[string]model.User
	mu    sync.RWMutex
}

func (s *Storage) Add(id string, u model.UserEntity) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.users[id] = model.User{UserEntity: u, ID: id}
	return nil
}

func NewStorage() *Storage {
	return &Storage{
		users: make(map[string]model.User, 0),
		mu:    sync.RWMutex{},
	}
}
