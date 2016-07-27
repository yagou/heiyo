package heiyo

import (
	"encoding/json"
	"fmt"
)

type Router struct {
	Route string      // 路由
	Body  interface{} // 包内容
}

// Handler defines route handler, middleware handler type.
type Handler func(context *Context, tag string)

var cacheRouter map[string][]Handler

func NewRouter(data []byte, tag string) error {
	rt := new(Router)
	err := json.Unmarshal(data, &rt)
	if err != nil {
		return err
	}
	fmt.Println(tag + "发消息了")
	fmt.Println(cacheRouter)
	fmt.Println(rt)
	fmt.Println(rt.Route)
	if rs, ok := cacheRouter[rt.Route]; ok {

		c := &Context{Body: rt.Body.(string)}
		for _, r := range rs {
			fmt.Println("开始发布信息")
			r(c, tag)
		}
	}
	return nil
}
