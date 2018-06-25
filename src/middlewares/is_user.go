package middlewares

import (
	"encoding/base64"
	"strings"

	"github.com/Electra-project/electra-api/src/helpers"
	"github.com/Electra-project/electra-api/src/libs/fail"
	"github.com/Electra-project/electra-api/src/libs/rpc"
	"github.com/Electra-project/electra-api/src/models"

	"github.com/gin-gonic/gin"
)

var userTokenModel = new(models.UserToken)

// IsUser checks the user Authorization header.
func IsUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if len(authorization) < 7 || !strings.HasPrefix(authorization, "Basic ") {
			fail.AnswerCustom(c, fail.Unauthorized, "")

			return
		}

		authorizationValue := helpers.Substring(authorization, 6)
		credentials, err := base64.StdEncoding.DecodeString(authorizationValue)
		if err != nil {
			fail.AnswerCustom(c, fail.Unauthorized, "")

			return
		}

		credentialsParts := strings.Split(string(credentials), ":")
		if len(credentialsParts) != 2 {
			fail.AnswerCustom(c, fail.Unauthorized, "")

			return
		}

		purseHash := credentialsParts[0]
		signature := credentialsParts[1]

		userToken, err := userTokenModel.GetByPurseHash(purseHash)
		if err != nil {
			fail.AnswerCustom(c, fail.Unauthorized, "")

			return
		}

		res, err := rpc.VerifyMessage(purseHash, signature, userToken.Challenge)
		if err != nil || !res.Result {
			fail.AnswerCustom(c, fail.Unauthorized, "")

			return
		}

		c.Set("purseHash", purseHash)
	}
}
