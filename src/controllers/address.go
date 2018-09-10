package controllers

import "github.com/gin-gonic/gin"
import "github.com/Electra-project/electra-api/src/models"

type AddressController struct{}

func (s AddressController) Get(c *gin.Context) {

	addrHash := c.Param("hash")

	addr, err := models.GetAddressFromDB(addrHash)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}
	c.Header("X-Version", "1.0")
	c.JSON(200, addr)

}
