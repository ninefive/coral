package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"github.com/ninefive/coral/models"
	"github.com/ninefive/coral/pkg/e"
	"github.com/ninefive/coral/pkg/util"
	"github.com/ninefive/coral/routers/api"
)

func Create(c *gin.Context) {
	log.Info("user create function called.", lager.Data{"X-Request-Id": util.GetRequestId(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		api.SendResponse(c, e.ErrBind, nil)
		return
	}

	u := models.User{
		Username: r.Username,
		Password: r.Password,
	}

	//validate the data.
	if err := u.Validate(); err != nil {
		api.SendResponse(c, e.ErrValidation, nil)
		return
	}

	//encrypt the user password.
	if err := u.Encrypt(); err != nil {
		api.SendResponse(c, e.ErrEncrypt, nil)
		return
	}

	//insert the user to the database.
	if err := u.Create(); err != nil {
		api.SendResponse(c, e.ErrDatabase, nil)
	}

	resp := CreateResponse{
		Username: r.Username,
	}

	//show user information.
	api.SendResponse(c, nil, resp)
}
