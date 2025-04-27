package storage

import (
	"errors"
	"sync"
)

type InMemStorage struct {
	mu sync.RWMutex
	users map[string]*User
}

func NewInMemStorage() *InMemStorage {
	return &InMemStorage{
		users: make(map[string]*User),
	}
}

func (s *InMemStorage) CreateUser(login string, hashedPassword []byte) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    if _, exists := s.users[login]; exists {
        return errors.New("user already exists")
    }
    
    s.users[login] = &User{
        Login:    login,
        Password: hashedPassword,
    }
    return nil
}

func (s *InMemStorage) GetUser(login string) (*User, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    user, exists := s.users[login]
    if !exists {
        return nil, errors.New("user not found")
    }
    return user, nil
}