package models

import "time"

// Merchant defines the struct of a Monzo merchant
type Merchant struct {
	ID              string            `json:"id"`
	Name            string            `json:"name"`
	Logo            string            `json:"logo"`
	Category        string            `json:"category"`
	Emoji           string            `json:"emoji"`
	GroupID         string            `json:"group_id"`
	Online          bool              `json:"online"`
	ATM             bool              `json:"atm"`
	Address         *Address          `json:"address"`
	Created         time.Time         `json:"created"`
	Metadata        map[string]string `json:"metadata"`
	DisableFeedback bool              `json:"disable_feedback"`
}
