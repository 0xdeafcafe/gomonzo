package models

import "time"

// Accounts defines the struct of a list of Monzo Accounts
type Accounts struct {
	Accounts []Account `json:"accounts"`
}

// Account defines the structure of a Monzo Account
type Account struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created"`
}
