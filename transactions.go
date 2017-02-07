package gomonzo

import (
	"fmt"

	"github.com/0xdeafcafe/gomonzo/helpers"
	"github.com/0xdeafcafe/gomonzo/models"
)

const (
	getTransactionURL   = "transactions/%s"
	listTransactionsURL = "transactions"
)

// ListTransactions gets all the transactions in an account
func (monzo *GoMonzo) ListTransactions(token *models.Token, accountID string) (*models.Transactions, *models.MonzoError, error) {
	params := map[string]string{
		"account_id": accountID,
		"expand[]":   "merchant",
	}

	resp, monzoError, err := monzo.httpHelper.Get(listTransactionsURL, params, token)
	if monzoError != nil || err != nil {
		return nil, monzoError, err
	}

	var transactions *models.Transactions
	err = helpers.UnmarshalJSON(resp.Body, &transactions)
	if err != nil {
		return nil, nil, err
	}

	return transactions, nil, nil
}

// GetTransaction gets a transaction by it's ID
func (monzo *GoMonzo) GetTransaction(token *models.Token, transactionID string) (*models.Transaction, *models.MonzoError, error) {
	params := map[string]string{
		"expand[]": "merchant",
	}

	url := fmt.Sprintf(getTransactionURL, transactionID)
	resp, monzoError, err := monzo.httpHelper.Get(url, params, token)
	if monzoError != nil || err != nil {
		return nil, monzoError, err
	}

	var container = &struct {
		Transaction *models.Transaction `json:"transaction"`
	}{}
	err = helpers.UnmarshalJSON(resp.Body, container)
	fmt.Println(container)
	if err != nil {
		return nil, nil, err
	}

	return container.Transaction, nil, nil
}
