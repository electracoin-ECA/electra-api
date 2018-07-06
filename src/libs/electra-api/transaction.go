package api

import (
	"errors"
	"fmt"

	"github.com/Electra-project/electra-api/src/helpers"
	"github.com/Electra-project/electra-api/src/models"
	mapset "github.com/deckarep/golang-set"
)

// GetTransactions fetches all the transactions based off of the block hash
func GetTransactions(blockHash string) ([]*models.Transaction, error) {
	var txns []*models.Transaction

	errOccured := helpers.GetJSON(fmt.Sprintf("%s/block/%s/txs", ApiURLPrefix, blockHash), &txns)
	if errOccured {
		return nil, errors.New(ErrMalformedResponseFromUpstreamAPI)
	}

	return txns, nil

}

// GetUniqueAddresses takes a variate number of transactions and returns all
// unique address present in them as strings
func GetUniqueAddresses(txns ...*models.Transaction) []string {
	addresses := mapset.NewSet()
	for _, txn := range txns {
		for _, vin := range txn.Vin {
			if vin.Address != CoinbaseAddress {
				addresses.Add(vin.Address)
			}
		}
		for _, vout := range txn.Vout {
			if vout.Address != CoinbaseAddress {
				addresses.Add(vout.Address)
			}
		}
	}
	return toStringSlice(addresses.ToSlice())
}

func toStringSlice(intfs []interface{}) []string {
	var strings []string
	for _, intf := range intfs {
		strings = append(strings, intf.(string))
	}
	return strings
}
