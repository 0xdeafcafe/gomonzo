package gomonzo

import (
	"github.com/0xdeafcafe/gomonzo/helpers"
	"github.com/0xdeafcafe/gomonzo/models"
)

const (
	listAccountsURL = "accounts"
)

// ListAccounts gets all the accounts belonging to the token
func (monzo *GoMonzo) ListAccounts(token *models.Token) (*models.Accounts, *models.MonzoError, error) {
	resp, monzoError, err := monzo.httpHelper.Get(listAccountsURL, nil, token)
	if monzoError != nil || err != nil {
		return nil, monzoError, err
	}

	var accounts *models.Accounts
	err = helpers.UnmarshalJSON(resp.Body, &accounts)
	if err != nil {
		return nil, nil, err
	}

	return accounts, nil, nil
}
