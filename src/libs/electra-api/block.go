package api

import (
	"context"
	"errors"
	"fmt"
	"log"

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

func GetPreviousBlocks(ctx context.Context, topBlockHash string) (chan BlockResponse, error) {
	blockChan := make(chan BlockResponse)

	blockresp, err := GetBlock(topBlockHash)

	if err != nil {
		return nil, err
	}

	go func(blockresp BlockResponse) {
		blockHash := blockresp.Block.Previousblockhash
		for {
			select {
			case <-ctx.Done():
				log.Println("received Cancel on the ctx")
				close(blockChan)
				return
			default:

				block, err := GetBlock(blockHash)
				if err == nil {
					blockChan <- *block
					blockHash = block.Block.Previousblockhash
				} else {
					close(blockChan)
					return
				}
			}

		}
	}(*blockresp)

	return blockChan, nil

}
