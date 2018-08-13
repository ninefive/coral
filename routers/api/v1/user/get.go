package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ninefive/coral/models"
	"github.com/ninefive/coral/pkg/e"
	"github.com/ninefive/coral/routers/api"
)

func Get(c *gin.Context) {
	username := c.Param("username")
	user, err := models.GetUser(username)
	if err != nil {
		api.SendResponse(c, e.ErrUserNotFound, nil)
		return
	}

	api.SendResponse(c, nil, user)
}
