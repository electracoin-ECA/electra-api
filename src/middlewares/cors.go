package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORS set the Access-Control headers.
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Accept, Accept-Encoding, Authorization, Content-Length, Content-Type, Origin")
		c.Header("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, POST, PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length")
		c.Header("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
	}
}
