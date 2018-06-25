package controllers

import (
	"net/http"

	"github.com/Electra-project/electra-api/src/models"

	"github.com/Electra-project/electra-api/src/libs/fail"
	"github.com/gin-gonic/gin"
)

// UserController describes a user controller.
type UserController struct{}

// Get retrieves the authenticated user data.
func (u UserController) Get(c *gin.Context) {
	purseHash := getPurseHash(c)

	user, err := userModel.GetByPurseHash(purseHash)
	if err != nil {
		fail.Answer(c, err, "user")

		return
	}

	c.Header("X-Version", "1.0")
	c.JSON(http.StatusOK, gin.H{"data": user})

	return
}

// Post creates a new database entry for the authenticated user.
func (u UserController) Post(c *gin.Context) {
	purseHash := getPurseHash(c)

	user, err := userModel.Insert(purseHash)
	if err != nil {
		fail.Answer(c, err, "user")

		return
	}

	c.Header("X-Version", "1.0")
	c.JSON(http.StatusCreated, gin.H{"data": user})
}

// Put update the authenticated user data.
func (u UserController) Put(c *gin.Context) {
	purseHash := getPurseHash(c)

	var data models.UserEditable
	err := c.BindJSON(&data)
	if err != nil {
		fail.Answer(c, err, "user")

		return
	}

	user, err := userModel.Update(purseHash, data)
	if err != nil {
		fail.Answer(c, err, "user")

		return
	}

	c.Header("X-Version", "1.0")
	c.JSON(http.StatusOK, gin.H{"data": user})
}
