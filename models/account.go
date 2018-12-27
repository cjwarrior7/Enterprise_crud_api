package models

// Account
type Account struct {
	ID      int64  `json:"id"`
	SuperUser    string `json:"is_superuser"`
}
