package controllers

import (
	"errors"

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
	txn, err := rpc.GetTransaction(txnID)
	if err != nil {
		fail.Answer(c, err, "transaction")
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
