package db

import (
	"errors"
	"webapp/model"
)

type DB interface {
	FindUser(username string) (*model.User, error)
	FindUserByUsername(username string) (*model.User, error)
	CreateUser(user model.User) error
	AppendFeedback(id string, feedback model.Feedback) error
}

var (
	ErrAlreadyExists = errors.New("user already exists")
	ErrNotFound      = errors.New("unknown user")
)

type db map[string]model.User

func New() DB {
	store := make(db)
	return store
}

func (s db) FindUser(username string) (*model.User, error) {
	if user, ok := s[username]; ok {
		return &user, nil
	}

	return nil, ErrNotFound
}

func (s db) FindUserByUsername(username string) (*model.User, error) {
	for _, v := range s {
		if v.Username == username {
			return &v, nil
		}
	}

	return nil, ErrNotFound
}

func (s db) CreateUser(user model.User) error {
	if _, ok := s[user.Username]; ok {
		return ErrAlreadyExists
	}
	s[user.Username] = user

	return nil
}

func (s db) AppendFeedback(username string, feedback model.Feedback) error {
	user, err := s.FindUser(username)
	if err != nil {
		return err
	}
	user.Feedback = append(user.Feedback, feedback)
	s[user.Username] = *user

	return nil
}
