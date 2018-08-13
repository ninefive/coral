package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ninefive/coral/models"
	"github.com/ninefive/coral/pkg/auth"
	"github.com/ninefive/coral/pkg/e"
	"github.com/ninefive/coral/pkg/token"
	"github.com/ninefive/coral/routers/api"
)

func Login(c *gin.Context) {
	var u models.User
	if err := c.Bind(&u); err != nil {
		api.SendResponse(c, e.ErrBind, nil)
		return
	}

	d, err := models.GetUser(u.Username)
	if err != nil {
		api.SendResponse(c, e.ErrUserNotFound, nil)
		return
	}

	if err := auth.Compare(d.Password, u.Password); err != nil {
		api.SendResponse(c, e.ErrPasswordIncorrect, nil)
		return
	}

	t, err := token.Sign(c, token.Context{ID: d.ID, Username: d.Username}, "")
	if err != nil {
		api.SendResponse(c, e.ErrToken, nil)
		return
	}

	api.SendResponse(c, nil, models.Token{Token: t})
}
