package models

import (
	"time"
)

// Transactions ..
type Transactions struct {
	Transactions []Transaction `json:"transactions"`
}

// Transaction defines the struct of a Monzo transaction
type Transaction struct {
	ID                string            `json:"id"`
	AccountID         string            `json:"account_id"`
	Description       string            `json:"description"`
	Merchant          *Merchant         `json:"merchant"`
	IsLoad            bool              `json:"is_load"`
	Currency          string            `json:"currency"`
	Amount            int64             `json:"amount"`
	AccountBalance    int64             `json:"account_balance"`
	Metadata          map[string]string `json:"metadata"`
	Notes             string            `json:"notes"`
	DeclineReason     string            `json:"decline_reason"`
	LocalAmount       int64             `json:"local_amount"`
	LocalCurrency     string            `json:"local_currency"`
	Scheme            string            `json:"scheme"`
	DedupeID          string            `json:"dedupe_id"`
	Originator        bool              `json:"originator"`
	IncludeInSpending bool              `json:"include_in_spending"`
	Created           time.Time         `json:"created"`
	Updated           *time.Time        `json:"updated"`
	Settled           string            `json:"settled"`
}
