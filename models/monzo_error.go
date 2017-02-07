package models

// MonzoError contains a Monzo API error
type MonzoError struct {
	Code string `json:"code"`

	Error            string `json:"invalid_request"`
	ErrorDescription string `json:"error_description"`

	Message string `json:"message"`
}
