package api

import (
	"errors"
	"fmt"

	"github.com/Electra-project/electra-api/src/models"

	"github.com/Electra-project/electra-api/src/helpers"
)

type LatestBlockResponse struct {
	Hash   string
	Height uint
}

type BlockResponse struct {
	Active        string
	Confirmations uint
	Block         *models.Block         `json:"block" bson:"block"`
	Txs           []*models.Transaction `json:"txs" bson:"txs"`
}

// GetLatestBlock gets the top most block from the chain
func GetLatestBlock() (*LatestBlockResponse, error) {
	var resp []LatestBlockResponse
	errOccured := helpers.GetJSON(fmt.Sprintf("%s/blocks_latest", ApiURLPrefix), &resp)

	if errOccured {
		return nil, errors.New(ErrMalformedResponseFromUpstreamAPI)
	}

	// for now not checking if the len > 0, as we should get blocks in all cases
	return &resp[0], nil
}

func GetBlock(blockHash string) (*BlockResponse, error) {
	var resp BlockResponse

	errOccured := helpers.GetJSON(fmt.Sprintf("%s/block/%s", ApiURLPrefix, blockHash), &resp)
	if errOccured {
		return nil, errors.New(ErrMalformedResponseFromUpstreamAPI)
	}

	return &resp, nil
}
