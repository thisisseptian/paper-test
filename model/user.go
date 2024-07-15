package model

type User struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Password []byte `json:"password"`
	WalletID int64  `json:"wallet_id"`
}

type UserDetail struct {
	User
	Wallet Wallet `json:"wallet"`
}
