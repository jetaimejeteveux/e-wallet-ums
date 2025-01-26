package models

type Wallet struct {
	ID      int `json:"id"`
	UserID  int `json:"user_id"`
	Balance int `json:"balance"`
}
