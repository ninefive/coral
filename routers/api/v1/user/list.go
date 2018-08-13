package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/ninefive/coral/pkg/e"
	"github.com/ninefive/coral/routers/api"

	"sync"

	"github.com/ninefive/coral/models"
	"github.com/ninefive/coral/pkg/util"
)

func List(c *gin.Context) {
	log.Info("list function called.")
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		api.SendResponse(c, e.ErrBind, nil)
		return
	}

	infos, count, err := ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	api.SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}

func ListUser(username string, offset, limit int) ([]*models.UserInfo, uint64, error) {
	infos := make([]*models.UserInfo, 0)
	users, count, err := models.ListUser(username, offset, limit)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, user := range users {
		ids = append(ids, user.ID)
	}

	wg := sync.WaitGroup{}
	userList := models.UserList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*models.UserInfo, len(users)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	for _, u := range users {
		wg.Add(1)
		go func(u *models.User) {
			defer wg.Done()

			shortId, err := util.GenShortID()
			if err != nil {
				errChan <- err
				return
			}

			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			userList.IdMap[u.ID] = &models.UserInfo{
				Id:        u.ID,
				Username:  u.Username,
				SayHello:  fmt.Sprintf("Hello %s", shortId),
				CreatedOn: u.CreatedOn.Format("2006-01-02 15:04:05"),
				UpdatedOn: u.UpdatedOn.Format("2006-01-02 15:04:05"),
			}
		}(u)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, count, err
	}

	for _, id := range ids {
		infos = append(infos, userList.IdMap[id])
	}

	return infos, count, nil
}
