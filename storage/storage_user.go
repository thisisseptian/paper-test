package storage

import (
	"errors"

	"paper-test/model"

	"golang.org/x/crypto/bcrypt"
)

var (
	userIDCounter int64
)

func incrementUserID() int64 {
	userIDCounter++
	return userIDCounter
}

func (s *Storage) GetUserByUsernameAndPassword(username, password string) (model.User, error) {
	for _, user := range s.users {
		if user.Username == username {
			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err == nil {
				return *user, nil
			}
		}
	}

	return model.User{}, errors.New("user not found or incorrect password")
}

func (s *Storage) GetUserByUserID(userID int64) (model.User, error) {
	user, exists := s.users[userID]
	if exists {
		return *user, nil
	}

	return model.User{}, errors.New("user data is not found")
}
