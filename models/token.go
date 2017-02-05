package models

import "time"

// Token holds the monzo response for creating or refreshing a token
type Token struct {
	AccessToken  string    `json:"access_token"`
	ClientID     string    `json:"client_id"`
	ExpiresIn    int       `json:"expires_in"`
	ExpiresAt    time.Time `json:"expires_at"`
	RefreshToken string    `json:"refresh_token"`
	TokenType    string    `json:"token_type"`
	UserID       string    `json:"user_id"`
}
