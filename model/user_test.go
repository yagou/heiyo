package model

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t.Log(time.Now())
}

func TestInsert(t *testing.T) {
	// user := userModel{
	// 	Username: "吴赐有s1",
	// 	Password: "Password",
	// 	Addtime:  time.Now(),
	// 	Status:   "Status",
	// }
	// // db.Create(&user)

	users := userModel{}
	usersSlice := []userModel{}
	rows, err := db.Model(users).Find(&usersSlice).Rows()
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	t.Log(usersSlice)
	t.Log(db)
}

func TestLogin(t *testing.T) {
	user := &userModel{
		Username: "吴赐有1",
		Password: "Password",
	}
	t.Log("登录验证")
	t.Log(user.Login())
	t.Log(user)

}
