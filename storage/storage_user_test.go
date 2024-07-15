package storage

import (
	"strconv"
	"testing"

	"paper-test/model"

	"golang.org/x/crypto/bcrypt"
)

func TestGetUserByUsernameAndPassword(t *testing.T) {
	store := &Storage{
		users: make(map[int64]*model.User),
	}

	password := "password123"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	userID := incrementUserID()
	store.users[userID] = &model.User{
		UserID:   userID,
		Username: "testuser",
		Password: hashedPassword,
		WalletID: 1,
	}

	tests := []struct {
		username string
		password string
		wantErr  bool
	}{
		{"testuser", "password123", false},
		{"testuser", "wrongpassword", true},
		{"wronguser", "password123", true},
	}

	for _, tt := range tests {
		t.Run(tt.username, func(t *testing.T) {
			user, err := store.GetUserByUsernameAndPassword(tt.username, tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByUsernameAndPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && user.Username != tt.username {
				t.Errorf("expected username %v, got %v", tt.username, user.Username)
			}
		})
	}
}

func TestGetUserByUserID(t *testing.T) {
	store := &Storage{
		users: make(map[int64]*model.User),
	}

	userID := incrementUserID()
	store.users[userID] = &model.User{
		UserID:   userID,
		Username: "testuser",
		Password: []byte("hashedpassword"),
		WalletID: 1,
	}

	tests := []struct {
		userID  int64
		wantErr bool
	}{
		{userID, false},
		{userID + 1, true},
	}

	for _, tt := range tests {
		t.Run(strconv.FormatInt(tt.userID, 10), func(t *testing.T) {
			user, err := store.GetUserByUserID(tt.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByUserID() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && user.UserID != tt.userID {
				t.Errorf("expected userID %v, got %v", tt.userID, user.UserID)
			}
		})
	}
}
