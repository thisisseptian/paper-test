package storage

import (
	"testing"

	"paper-test/model"
)

func TestGetWalletByWalletID(t *testing.T) {
	store := &Storage{
		wallets: make(map[int64]*model.Wallet),
	}

	wallet := model.Wallet{
		WalletID: 1,
		Balance:  10000000,
	}

	store.wallets[wallet.WalletID] = &wallet

	retrievedWallet, err := store.GetWalletByWalletID(wallet.WalletID)
	if err != nil {
		t.Fatalf("failed to get wallet: %v", err)
	}

	if retrievedWallet.WalletID != wallet.WalletID {
		t.Errorf("expected wallet ID %v, got %v", wallet.WalletID, retrievedWallet.WalletID)
	}

	if retrievedWallet.Balance != wallet.Balance {
		t.Errorf("expected wallet balance %v, got %v", wallet.Balance, retrievedWallet.Balance)
	}
}

func TestGetWalletByWalletID_NotFound(t *testing.T) {
	store := &Storage{
		wallets: make(map[int64]*model.Wallet),
	}

	_, err := store.GetWalletByWalletID(1)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	expectedErr := "wallet data is not found"
	if err.Error() != expectedErr {
		t.Errorf("expected error message %v, got %v", expectedErr, err.Error())
	}
}

func TestUpdateWalletBalanceByWalletID(t *testing.T) {
	store := &Storage{
		wallets: make(map[int64]*model.Wallet),
	}

	wallet := model.Wallet{
		WalletID: 1,
		Balance:  10000000,
	}

	store.wallets[wallet.WalletID] = &wallet

	updatedWallet := model.Wallet{
		WalletID: 1,
		Balance:  20000000,
	}

	err := store.UpdateWalletBalanceByWalletID(updatedWallet)
	if err != nil {
		t.Fatalf("failed to update wallet balance: %v", err)
	}

	retrievedWallet, err := store.GetWalletByWalletID(wallet.WalletID)
	if err != nil {
		t.Fatalf("failed to get wallet: %v", err)
	}

	if retrievedWallet.Balance != updatedWallet.Balance {
		t.Errorf("expected wallet balance %v, got %v", updatedWallet.Balance, retrievedWallet.Balance)
	}
}

func TestUpdateWalletBalanceByWalletID_NotFound(t *testing.T) {
	store := &Storage{
		wallets: make(map[int64]*model.Wallet),
	}

	wallet := model.Wallet{
		WalletID: 1,
		Balance:  20000000,
	}

	err := store.UpdateWalletBalanceByWalletID(wallet)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	expectedErr := "fail update user amount"
	if err.Error() != expectedErr {
		t.Errorf("expected error message %v, got %v", expectedErr, err.Error())
	}
}
