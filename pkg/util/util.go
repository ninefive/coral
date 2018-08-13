package util

import (
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

func GenShortID() (string, error) {
	return shortid.Generate()
}

func GetRequestId(ctx *gin.Context) string {
	v, ok := ctx.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}
