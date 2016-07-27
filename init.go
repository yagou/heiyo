package heiyo

import (
	"fmt"
	"github.com/yagou/heiyo/packet"
	"github.com/yagou/heiyo/router"
)

func init() {

	cacheRouter = make(map[string][]Handler)

	// 聊天
	cacheRouter[router.ONE_TO_ONE] = []Handler{chat}

	// 登录认证
	cacheRouter[router.LOGIN] = []Handler{login}

	cacheRouter[router.ONE_TO_MANY] = []Handler{chatByONE_TO_MANY}
	cacheRouter[router.GET_USER_LIST] = []Handler{GET_USER_LIST}

}

//  登录验证
func login(context *Context, tag string) {

}

// 一对一聊天
func chat(context *Context, tag string) {
	fmt.Printf("tag:%s ", tag)
	fmt.Println(context)
	conn, err := hy_conn_cache.Get_conn(tag)
	if err == nil {
		conn.Write(packet.Packet([]byte("你好啊")))
	}
}

// 一对多
func chatByONE_TO_MANY(context *Context, tag string) {

	fmt.Printf("tag:%s ", tag)

	conns := hy_conn_cache.Get_conn_all()

	for k, conn := range conns {
		if tag != k {
			conn.Write(packet.Packet([]byte(tag + "对所有人说：" + context.Body)))
		}
	}
}

//  获取在线用户列表
func GET_USER_LIST(context *Context, tag string) {
	conns := hy_conn_cache.Get_conn_all()

	var userList string
	count := 0
	for k, _ := range conns {
		if tag != k {
			userList += fmt.Sprintf("\n %d、tag:%s ", count, k)
			count++
		}
	}
	conn, err := hy_conn_cache.Get_conn(tag)
	fmt.Println("获取用户列表：", userList)
	if err == nil {
		fmt.Println("开始发送用户列表")
		conn.Write(packet.Packet([]byte(userList)))
	}
}
