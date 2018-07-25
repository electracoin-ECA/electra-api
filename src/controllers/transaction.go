package controllers

import (
	"github.com/Electra-project/electra-api/src/libs/rpc"
	"github.com/gin-gonic/gin"
)

type TransactionController struct{}

// Get a user public data.
func (s TransactionController) Get(c *gin.Context) {

	txnID := c.Param("id")
	txn, err := rpc.GetTransaction(txnID)
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error."})
		c.Abort()

		return
	}

	c.Header("X-Version", "1.0")
	c.JSON(200, txn)

	return
}

func (s TransactionController) Post(c *gin.Context) {
}
