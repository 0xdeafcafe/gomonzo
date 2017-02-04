package models

// MonzoError contains a Monzo API error
type MonzoError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	Message          string `json:"message"`
}
