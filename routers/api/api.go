package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ninefive/coral/pkg/e"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, msg := e.DecodeErr(err)

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}
