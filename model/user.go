package model

import (
	"strings"
	"time"
)

type userModel struct {
	Id       uint `gorm:"primary_key"`
	Username string
	Password string
	Addtime  time.Time
	Status   string
}

func (u *userModel) TableName() string {
	return "hy_user"
}

// 用户登录
func (u *userModel) Login() bool {

	var userRows []userModel
	rows, err := db.Model(u).Where(userModel{Username: u.Username}).Find(&userRows).Rows()
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	for k, _ := range userRows {
		if strings.EqualFold(userRows[k].Password, u.Password) {
			return true
		}
	}
	return false
}
