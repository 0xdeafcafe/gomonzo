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
	url := fmt.Sprintf(getTransactionURL, transactionID)
	resp, monzoError, err := monzo.httpHelper.Get(url, nil, token)
	if monzoError != nil || err != nil {
		return nil, monzoError, err
	}

	var transaction *models.Transaction
	err = helpers.UnmarshalJSON(resp.Body, &transaction)
	if err != nil {
		return nil, nil, err
	}

	return transaction, nil, nil
}
