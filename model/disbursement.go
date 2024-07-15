package model

import "time"

type Disbursement struct {
	DisbursementID int64       `json:"disbursement_id"`
	FromWalletID   int64       `json:"from_wallet_id"`
	ToBankAccount  BankAccount `json:"to_bank_account"`
	Amount         float64     `json:"amount"`
	TimeStamp      time.Time   `json:"time_stamp"`
}

type BankAccount struct {
	BankAccountNumber string `json:"bank_account_number"`
	BankAccountName   string `json:"bank_account_name"`
	BankName          string `json:"bank_name"`
}

type DisbursementRequest struct {
	Amount            float64 `json:"amount"`
	BankAccountNumber string  `json:"bank_account_number"`
	BankAccountName   string  `json:"bank_account_name"`
	BankName          string  `json:"bank_name"`
}
