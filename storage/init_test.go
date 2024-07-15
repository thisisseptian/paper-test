package storage

import (
	"paper-test/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestNewStorage(t *testing.T) {
	storage := NewStorage()

	assert.NotNil(t, storage)
}

func TestInitData(t *testing.T) {
	store := &Storage{
		users:   make(map[int64]*model.User),
		wallets: make(map[int64]*model.Wallet),
	}

	store.InitData()

	if len(store.wallets) != 1 {
		t.Fatalf("expected 1 wallet, got %d", len(store.wallets))
	}

	var wallet *model.Wallet
	for _, w := range store.wallets {
		wallet = w
		break
	}

	if wallet.Balance != 10000000 {
		t.Errorf("expected wallet balance 10000000, got %v", wallet.Balance)
	}

	if len(store.users) != 1 {
		t.Fatalf("expected 1 user, got %d", len(store.users))
	}

	var user *model.User
	for _, u := range store.users {
		user = u
		break
	}

	if user.Username != "Septian" {
		t.Errorf("expected username Septian, got %v", user.Username)
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte("password123")); err != nil {
		t.Errorf("password does not match, error: %v", err)
	}

	if user.WalletID != wallet.WalletID {
		t.Errorf("expected wallet ID %v, got %v", wallet.WalletID, user.WalletID)
	}
}
