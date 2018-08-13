package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ninefive/coral/models"
	"github.com/ninefive/coral/pkg/e"
	"github.com/ninefive/coral/routers/api"
)

func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := models.DeleteUserById(uint64(userId)); err != nil {
		api.SendResponse(c, e.ErrDatabase, nil)
		return
	}

	api.SendResponse(c, nil, nil)
}
