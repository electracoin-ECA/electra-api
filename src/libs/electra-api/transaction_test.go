package api

import (
	"errors"
	"testing"

	"github.com/Electra-project/electra-api/src/models"
)

func TestGetTransaction(t *testing.T) {

	transactions := getLatestTransaction(t)
	if len(transactions) == 0 {
		t.Fatal(errors.New("The transactions returned are empty"))
	}
}

func TestGetUniqueAddresses(t *testing.T) {
	transactions := getLatestTransaction(t)

	addresses := GetUniqueAddresses(transactions...)

	if len(addresses) == 0 {
		t.Fatal(errors.New("there are no addresses in the transactions"))
	}

	// t.Logf("%+v", addresses)
}

func getLatestTransaction(t *testing.T) []*models.Transaction {
	resp, err := GetLatestBlock()
	if err != nil {
		t.Fatal(err)
	}

	transactions, err := GetTransactions(resp.Hash)

	if err != nil {
		t.Fatal(err)
	}

	return transactions
}
