package fail

import (
	"net/http"
	"strings"

	"github.com/Electra-project/electra-api/src/helpers"
	"github.com/gin-gonic/gin"
)

const (
	// MissingParameter error represents a missing parameter in the request query.
	MissingParameter = iota
	// MissingProperty error represents a missing property in the request body.
	MissingProperty
	// NotFound error represents an unfound entity in the database.
	NotFound
	// Unauthorized error represents a wrong authentication header.
	Unauthorized
	// WrongPropertyValue error represents an unexpected property value.
	WrongPropertyValue
)

// Answer sends a JSON response error.
func Answer(c *gin.Context, err error, entityLabel string) {
	errMessage := err.Error()
	helpers.LogErr(errMessage)

	switch true {

	case errMessage == "not found":
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{"message": strings.Title(entityLabel) + " not found."},
		)

	case strings.Contains(errMessage, "invalid character"):
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			gin.H{"message": "Malformed body JSON."},
		)

	case strings.Contains(errMessage, "cannot unmarshal"):
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			gin.H{"message": "Wrong body properties types."},
		)

	case strings.Contains(errMessage, "duplicate key error"):
		c.AbortWithStatusJSON(
			http.StatusConflict,
			gin.H{"message": "This " + entityLabel + " already exists."},
		)

	default:
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"message": "Internal server error."},
		)
	}
}

// AnswerCustom sends a JSON response error related to a constant-defined error.
func AnswerCustom(c *gin.Context, errorIndex uint8, target string) {
	switch errorIndex {

	case MissingParameter:
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			gin.H{"message": "Missing query parameter: " + target + "."},
		)

	case MissingProperty:
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			gin.H{"message": "Missing body property: " + target + "."},
		)

	case NotFound:
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{"message": strings.Title(target) + " not found."},
		)

	case Unauthorized:
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"message": "Non-provided or wrong credentials."},
		)

	case WrongPropertyValue:
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			gin.H{"message": "Wrong body property value for: " + target + "."},
		)

	default:
		helpers.LogErr(`Error: Couldn't handle this custom error.`)

		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"message": "Internal server error."},
		)
	}
}
