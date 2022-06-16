package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CatchError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  err.(string),
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
