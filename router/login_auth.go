package router

import (
	"github.com/yagou/heiyo/conn"
)

func login_auth(context *Context, tag string) {

	_, err := conn.NewHYConn().Get_conn(tag)
	if err != nil {
		conn.NewHYConn().Del_conn(tag)
	}

}
