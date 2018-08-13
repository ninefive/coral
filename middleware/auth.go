package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ninefive/coral/pkg/e"
	"github.com/ninefive/coral/pkg/token"
	"github.com/ninefive/coral/routers/api"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			api.SendResponse(c, e.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
