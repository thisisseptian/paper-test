package model

type Wallet struct {
	WalletID int64   `json:"wallet_id"`
	Balance  float64 `json:"balance"`
}
