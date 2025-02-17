package storage

import (
	"errors"
	"sync"
)

type User struct {
	Id      int
	Name    string
	CardPin int
}

type Storage interface {
	addUser(user User) error
	getUser(Id int) (User, error)
	updateUser(Id int, updated User) error
	deleteUser(Id int) error
}

type ReadyStorage struct {
	users map[int]User
	mutex sync.RWMutex
}

func CreateStorage() *ReadyStorage {
	return &ReadyStorage{
		users: make(map[int]User),
	}
}

func (s *ReadyStorage) AddUser(user User) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.users[user.Id]; ok {
		return errors.New("user already exists")
	}

	s.users[user.Id] = user
	return nil
}

func (s *ReadyStorage) GetUser(Id int) (User, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	user, ok := s.users[Id]
	if !ok {
		return User{}, errors.New("item not found")
	}

	return user, nil

}

func (s *ReadyStorage) deleteUser(Id int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.users[Id]; !ok {
		return errors.New("item not found")
	}
	delete(s.users, Id)
	return nil
}

func (s *ReadyStorage) updateUser(Id int, updated User) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.users[Id]; !ok {
		return errors.New("item not found")
	}

	s.users[Id] = updated
	return nil
}
