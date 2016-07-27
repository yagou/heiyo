package router

import (
	"github.com/yagou/heiyo/conn"
)

/**
 * 群聊信息
 */
func group_chat(context *Context, tag string) {

	conns := conn.NewHYConn().Get_conn_all()

	for k, conn := range conns {
		if tag != k {
			// 群聊，将自己挂除在外
		}
		conn.Write([]byte(tag + "对所有人说：" + context.Body))
	}
}
