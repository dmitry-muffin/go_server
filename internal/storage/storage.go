package storage

import (
	"errors"
	"sync"
)

type User struct {
	id       int
	name     string
	card_pin int
}

type Storage interface {
	addUser(user User) error
	getUser(id int) (User, error)
	updateUser(id int, updated User) error
	deleteUser(id int) error
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

func (s *ReadyStorage) addUser(user User) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.users[user.id]; ok {
		return errors.New("user already exists")
	}

	s.users[user.id] = user
	return nil
}

func (s *ReadyStorage) getUser(id int) (User, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	user, ok := s.users[id]
	if !ok {
		return User{}, errors.New("item not found")
	}

	return user, nil

}

func (s *ReadyStorage) deleteUser(id int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.users[id]; !ok {
		return errors.New("item not found")
	}
	delete(s.users, id)
	return nil
}

func (s *ReadyStorage) updateUser(id int, updated User) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.users[id]; !ok {
		return errors.New("item not found")
	}

	s.users[id] = updated
	return nil
}
