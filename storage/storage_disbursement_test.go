package storage

import (
	"sort"
	"testing"
	"time"

	"paper-test/model"
)

func TestInsertDisbursement(t *testing.T) {
	store := &Storage{
		disbursements: make(map[int64]*model.Disbursement),
	}

	disbursement := model.Disbursement{
		FromWalletID: 1,
		ToBankAccount: model.BankAccount{
			BankAccountNumber: "123456789",
			BankAccountName:   "Septian",
			BankName:          "BCA",
		},
		Amount:    10000000,
		TimeStamp: time.Now(),
	}

	disbursementID, err := store.InsertDisbursement(disbursement)
	if err != nil {
		t.Fatalf("failed to insert disbursement: %v", err)
	}

	if _, exists := store.disbursements[disbursementID]; !exists {
		t.Fatalf("disbursement with ID %d does not exist in storage", disbursementID)
	}

	if store.disbursements[disbursementID].Amount != disbursement.Amount {
		t.Errorf("expected disbursement amount %v, got %v", disbursement.Amount, store.disbursements[disbursementID].Amount)
	}
}

func TestGetDisbursements(t *testing.T) {
	store := &Storage{
		disbursements: make(map[int64]*model.Disbursement),
	}

	walletID := int64(1)
	disbursements := []model.Disbursement{
		{
			FromWalletID: walletID,
			ToBankAccount: model.BankAccount{
				BankAccountNumber: "123456789",
				BankAccountName:   "Septian",
				BankName:          "BCA",
			},
			Amount:    10000000,
			TimeStamp: time.Now(),
		},
		{
			FromWalletID: walletID,
			ToBankAccount: model.BankAccount{
				BankAccountNumber: "987654321",
				BankAccountName:   "Pratama",
				BankName:          "BNI",
			},
			Amount:    20000000,
			TimeStamp: time.Now(),
		},
		{
			FromWalletID: 2,
			ToBankAccount: model.BankAccount{
				BankAccountNumber: "555555555",
				BankAccountName:   "Rusmana",
				BankName:          "Mandiri",
			},
			Amount:    30000000,
			TimeStamp: time.Now(),
		},
	}

	for _, disbursement := range disbursements {
		_, err := store.InsertDisbursement(disbursement)
		if err != nil {
			t.Fatalf("failed to insert disbursement: %v", err)
		}
	}

	retrievedDisbursements, err := store.GetDisbursements(walletID)
	if err != nil {
		t.Fatalf("failed to get disbursements: %v", err)
	}

	expectedCount := 2
	if len(retrievedDisbursements) != expectedCount {
		t.Errorf("expected %d disbursements, got %d", expectedCount, len(retrievedDisbursements))
	}

	expectedDisbursements := disbursements[:2]
	sort.Slice(expectedDisbursements, func(i, j int) bool {
		return expectedDisbursements[i].DisbursementID < expectedDisbursements[j].DisbursementID
	})

	for i, disbursement := range retrievedDisbursements {
		if disbursement.Amount != expectedDisbursements[i].Amount {
			t.Errorf("expected disbursement amount %v, got %v", expectedDisbursements[i].Amount, disbursement.Amount)
		}
		if disbursement.ToBankAccount.BankAccountNumber != expectedDisbursements[i].ToBankAccount.BankAccountNumber {
			t.Errorf("expected bank account number %v, got %v", expectedDisbursements[i].ToBankAccount.BankAccountNumber, disbursement.ToBankAccount.BankAccountNumber)
		}
		if disbursement.ToBankAccount.BankAccountName != expectedDisbursements[i].ToBankAccount.BankAccountName {
			t.Errorf("expected bank account name %v, got %v", expectedDisbursements[i].ToBankAccount.BankAccountName, disbursement.ToBankAccount.BankAccountName)
		}
		if disbursement.ToBankAccount.BankName != expectedDisbursements[i].ToBankAccount.BankName {
			t.Errorf("expected bank name %v, got %v", expectedDisbursements[i].ToBankAccount.BankName, disbursement.ToBankAccount.BankName)
		}
	}
}
