package controllers

import (
	"github.com/Electra-project/electra-api/src/libs/rpc"
	"github.com/gin-gonic/gin"
)

// StatusController describes a status controller.
type StatusController struct{}

// Get a user public data.
func (s StatusController) Get(c *gin.Context) {
	status, err := rpc.GetInfo()
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error."})
		c.Abort()

		return
	}

	c.Header("X-Version", "1.0")
	c.JSON(200, status)

	return
}
