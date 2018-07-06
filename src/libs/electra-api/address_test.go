package api

import (
	"errors"
	"testing"
)

func TestAddressDatas(t *testing.T) {
	transactions := getLatestTransaction(t)

	addresses := GetUniqueAddresses(transactions...)

	addressChan := GetAddressDatas(addresses...)

	var fin int
	for address := range addressChan {
		fin++
		t.Log("Address: ", address.Addr, "Balance:", address.Balance)
		if address.Addr == "" {
			t.Fatal(errors.New("One of the address received is blank"))
		}
	}

	if len(addresses) != fin {
		t.Fatal("Some address were not found on the server but were present in the block")
	}
}
