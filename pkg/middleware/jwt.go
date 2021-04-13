package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// A jwt middleware to check tokens
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = http.StatusOK
		token := c.Query("token")
		if token == "" {
			code = http.StatusPreconditionFailed
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = http.StatusRequestTimeout
				default:
					code = http.StatusNetworkAuthenticationRequired
				}
			}
		}

		if code != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"data": data,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
