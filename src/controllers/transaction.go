package controllers

import (
	"errors"

	"encoding/json"

	"github.com/Electra-project/electra-api/src/libs/fail"
	"github.com/Electra-project/electra-api/src/libs/rpc"
	"github.com/gin-gonic/gin"
)

type TransactionRequest struct {
	Hash string
}

type TransactionController struct{}

// Get a user public data.
func (s TransactionController) Get(c *gin.Context) {

	txnID := c.Param("id")

	body, resp, err := rpc.QueryBytesAndResp("gettransaction", []string{
		txnID,
	})
	if err != nil {
		fail.Answer(c, err, "transaction")
		return
	}

	if resp.StatusCode == 500 {
		errResp := make(map[string]interface{})
		err := json.Unmarshal(body, &errResp)
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal Error"})
			return
		}
		c.JSON(500, errResp)
		return
	}

	var txn rpc.GetTransactionResponse
	err = json.Unmarshal(body, &txn)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Error"})
		return
	}

	c.Header("X-Version", "1.0")
	c.JSON(200, gin.H{"txn": txn})
	return
}

func (s TransactionController) Post(c *gin.Context) {
	txnReq := TransactionRequest{}
	err := c.BindJSON(&txnReq)

	if err != nil {
		fail.Answer(c, err, "transaction")
		return
	}
	txnHash := rpc.SendTransaction(txnReq.Hash)
	if txnHash == "" {
		fail.Answer(c, errors.New("Blank txn hash returned"), "transaction")
	}

	c.Header("X-Version", "1.0")
	c.JSON(200, gin.H{"hash": txnHash})

	return

}
