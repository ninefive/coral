package models

import (
	"fmt"

	"github.com/ninefive/coral/pkg/auth"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	BaseModel
	Username string `json:"username" gorm:"column:username; not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password; not null" binding:"required" validate:"min=5,max=128"`
}

func (user *User) TableName() string {
	return "t_users"
}

//add user
func (user *User) Create() error {
	return DB.Self.Create(&user).Error
}

//delete user by id
func DeleteUserById(id uint64) error {
	user := User{}
	user.BaseModel.ID = id
	return DB.Self.Delete(&user).Error
}

//update user
func (user *User) Update() error {
	return DB.Self.Save(user).Error
}

//get user by username
func GetUser(username string) (*User, error) {
	u := &User{}
	d := DB.Self.Where("username = ?", username).First(&u)
	return u, d.Error
}

//list users
func ListUser(username string, offset, limit int) ([]*User, uint64, error) {
	if limit == 0 {
		limit = viper.GetInt("page_size")
	}

	users := make([]*User, 0)
	var count uint64

	where := fmt.Sprintf("username like '%%%s%%'", username)
	if err := DB.Self.Model(&User{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

func (u *User) Compare(password string) (err error) {
	err = auth.Compare(u.Password, password)
	return
}

func (u *User) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

func (u *User) Validate() error {
	valid := validator.New()
	return valid.Struct(u)
}
