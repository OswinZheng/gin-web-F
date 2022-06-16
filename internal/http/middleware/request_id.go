package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuidInfo, _ := uuid.NewV4()
		c.Writer.Header().Set("X-Request-Id", uuidInfo.String())
		c.Next()
	}
}
