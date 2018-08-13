package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"github.com/ninefive/coral/models"
	"github.com/ninefive/coral/pkg/e"
	"github.com/ninefive/coral/pkg/util"
	"github.com/ninefive/coral/routers/api"
)

func Update(c *gin.Context) {
	log.Info("update function called.", lager.Data{"X-Request-Id": util.GetRequestId(c)})

	userId, _ := strconv.Atoi(c.Param("id"))

	var u models.User
	if err := c.Bind(&u); err != nil {
		api.SendResponse(c, e.ErrBind, nil)
		return
	}

	u.ID = uint64(userId)

	if err := u.Validate(); err != nil {
		api.SendResponse(c, e.ErrValidation, nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		api.SendResponse(c, e.ErrEncrypt, nil)
		return
	}

	if err := u.Update(); err != nil {
		api.SendResponse(c, e.ErrDatabase, nil)
		return
	}

	api.SendResponse(c, nil, nil)
}
