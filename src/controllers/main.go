package controllers

import (
	"github.com/Electra-project/electra-api/src/libs/fail"
	"github.com/Electra-project/electra-api/src/models"
	"github.com/gin-gonic/gin"
)

var userModel = new(models.User)
var userTokenModel = new(models.UserToken)

func getPurseHash(c *gin.Context) string {
	purseHashGeneric, is := c.Get("purseHash")
	if is == false {
		fail.AnswerCustom(c, fail.Unauthorized, "")

		return ""
	}

	return purseHashGeneric.(string)
}
