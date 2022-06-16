package middleware

import (
	"net/http"

	"github.com/OswinZheng/gin-web-F/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var msg string
		var code = http.StatusOK
		token := getToken(c)
		if token == "" {
			code = http.StatusUnauthorized
			msg = "token is empty"
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				code = http.StatusUnauthorized
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					msg = "token is expired"
				default:
					msg = "token is invalid"
				}
			}
		}

		if code != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  msg,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func getToken(c *gin.Context) string {
	token := c.GetHeader("Authorization")
	if token == "" {
		token = c.Query("token")
	}

	return token
}
