package controllers

import (
	"net/http"

	"github.com/Electra-project/electra-auth/src/libs/fail"
	"github.com/gin-gonic/gin"
)

// UserTokenController class.
type UserTokenController struct{}

// Get a user token challenge from a Purse Account address hash.
func (u UserTokenController) Get(c *gin.Context) {
	purseHash := c.Param("purseHash")

	if !isPurseHashValid(purseHash) {
		fail.AnswerCustom(c, fail.WrongPropertyValue, "purseHash")

		return
	}

	userToken, err := userTokenModel.GetByPurseHash(c.Param("purseHash"))
	if err != nil {
		fail.Answer(c, err, "user token")

		return
	}

	c.Header("X-Version", "1.0")
	c.JSON(http.StatusOK, gin.H{"data": userToken})

	return
}

// Post a user token challenge from a Purse Account address hash.
func (u UserTokenController) Post(c *gin.Context) {
	purseHash := c.Param("purseHash")

	if !isPurseHashValid(purseHash) {
		fail.AnswerCustom(c, fail.WrongPropertyValue, "purseHash")

		return
	}

	userToken, err := userTokenModel.Insert(c.Param("purseHash"))
	if err != nil {
		fail.Answer(c, err, "user token")

		return
	}

	c.Header("X-Version", "1.0")
	c.JSON(http.StatusCreated, gin.H{"data": userToken})

	return
}

func isPurseHashValid(purseHash string) bool {
	return len(purseHash) == 34 && string([]rune(purseHash)[0]) == "E"
}
