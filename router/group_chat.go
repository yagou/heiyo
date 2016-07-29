package router

import (
	"encoding/json"
	"github.com/yagou/heiyo/conn"
)

/**
 * 群聊信息
 */
func group_chat(context *Context, tag string) {

	conns := conn.NewHYConn().Get_conn_all()
	ct := new(Context)
	ct.Body = context.Body
	ct.Target = "chat"

	for k, conn := range conns {
		if tag != k {
			// 群聊，将自己挂除在外
			ct.Event = "other"
		} else {
			ct.Event = "self"
		}

		body, err := json.Marshal(ct)
		if err != nil {
			continue
		}
		conn.Write(body)
	}
}
