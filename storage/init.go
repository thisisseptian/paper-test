package storage

import (
	"paper-test/model"

	"golang.org/x/crypto/bcrypt"
)

type IStorage interface {
	InitData()
	GetUserByUsernameAndPassword(username, password string) (model.User, error)
	GetUserByUserID(userID int64) (model.User, error)
	GetWalletByWalletID(walletID int64) (model.Wallet, error)
	InsertDisbursement(disbursement model.Disbursement) (int64, error)
	UpdateWalletBalanceByWalletID(wallet model.Wallet) error
	GetDisbursements(walletID int64) ([]model.Disbursement, error)
}

type Storage struct {
	users         map[int64]*model.User
	wallets       map[int64]*model.Wallet
	disbursements map[int64]*model.Disbursement
	IStorage
}

func NewStorage() *Storage {
	return &Storage{
		users:         make(map[int64]*model.User),
		wallets:       make(map[int64]*model.Wallet),
		disbursements: make(map[int64]*model.Disbursement),
	}
}

func (s *Storage) InitData() {
	// initialize some data for testing purposes
	// create wallet
	walletID := incrementWalletID()
	wallet := &model.Wallet{
		WalletID: walletID,
		Balance:  10000000,
	}

	s.wallets[walletID] = wallet

	// create user
	userID := incrementUserID()
	password, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	user := &model.User{
		UserID:   userID,
		Username: "Septian",
		Password: password,
		WalletID: walletID,
	}

	s.users[userID] = user
}
