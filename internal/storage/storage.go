package storage

import (
	"fmt"
	"sync"

	"go_server/internal/domain"
)

type Storage interface {
	addUser(user domain.User) error
	getUser(Id int) (domain.User, error)
	updateUser(Id int, updated domain.User) error
	deleteUser(Id int) error
}

type ReadyStorage struct {
	users map[int]domain.User
	mutex *sync.RWMutex
}

func CreateStorage() *ReadyStorage {
	return &ReadyStorage{
		users: make(map[int]domain.User),
	}
}

func (s *ReadyStorage) AddUser(user domain.User) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.users[user.Id]; ok {
		return fmt.Errorf("user already exists")
	}

	s.users[user.Id] = user
	return nil
}

func (s *ReadyStorage) GetUser(Id int) (domain.User, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	user, ok := s.users[Id]
	if !ok {
		return domain.User{}, fmt.Errorf("item not found")
	}

	return user, nil

}

func (s *ReadyStorage) deleteUser(Id int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.users[Id]; !ok {
		return fmt.Errorf("item not found")
	}
	delete(s.users, Id)
	return nil
}

func (s *ReadyStorage) updateUser(Id int, updated domain.User) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.users[Id]; !ok {
		return fmt.Errorf("item not found")
	}

	s.users[Id] = updated
	return nil
}
