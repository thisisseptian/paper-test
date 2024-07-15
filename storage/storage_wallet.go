package storage

import (
	"errors"

	"paper-test/model"
)

var (
	walletIDCounter int64
)

func incrementWalletID() int64 {
	walletIDCounter++
	return walletIDCounter
}

func (s *Storage) GetWalletByWalletID(walletID int64) (model.Wallet, error) {
	wallet, exists := s.wallets[walletID]
	if exists {
		return *wallet, nil
	}

	return model.Wallet{}, errors.New("wallet data is not found")
}

func (s *Storage) UpdateWalletBalanceByWalletID(wallet model.Wallet) error {
	data, exists := s.wallets[wallet.WalletID]
	if !exists {
		return errors.New("fail update user amount")
	}

	data.Balance = wallet.Balance

	return nil
}
