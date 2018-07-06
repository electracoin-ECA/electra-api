package api

import (
	"errors"
	"fmt"

	"github.com/Electra-project/electra-api/src/helpers"
	"github.com/Electra-project/electra-api/src/models"
)

// CoinbaseAddress is a constant which tells us that the address in the transaction is
// done using coinbase. We want to avoid adding this to the database
const CoinbaseAddress = "coinbase"

const MaxReq = 5

// GetAddressData gets the balance and other related information for the provided address
func GetAddressData(address string) (*models.Address, error) {

	var addr models.Address
	errOccured := helpers.GetJSON(fmt.Sprintf("%s/getaddress/%s", ApiURLPrefix, address), &addr)

	if errOccured {
		return nil, errors.New(ErrMalformedResponseFromUpstreamAPI)
	}

	return &addr, nil

}

// GetAddressDatas gets data related to addresses in the form of a channel
func GetAddressDatas(addresses ...string) chan models.Address {
	controlChan := make(chan bool, MaxReq)
	defer close(controlChan)
	controlChan = fillChannel(controlChan, MaxReq)
	addressChan := make(chan models.Address, len(addresses))
	go func() {
		for _, address := range addresses {
			<-controlChan
			go func(address string) {
				addr, err := GetAddressData(address)
				if err == nil {
					addressChan <- *addr
				}
				controlChan <- true
			}(address)
		}
	}()

	return addressChan

}

func fillChannel(c chan bool, factor int) chan bool {
	for _i := 0; _i < factor; _i++ {
		c <- true // doesn't matter if true or false
	}
	return c
}
