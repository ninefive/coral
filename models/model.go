package models

import (
	"sync"
	"time"
)

type BaseModel struct {
	ID        uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreatedOn time.Time  `gorm:"column:created_on" json:"createdOn"`
	UpdatedOn time.Time  `gorm:"column:updated_on" json:"updatedOn"`
	DeletedOn *time.Time `gorm:"column:deleted_on" sql:"index" json:"deleted_on"`
}

type UserInfo struct {
	Id        uint64 `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	SayHello  string `json:"sayHello"`
	CreatedOn string `json:"created_on"`
	UpdatedOn string `json:"updated_on"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}

type Token struct {
	Token string `json:"token"`
}
