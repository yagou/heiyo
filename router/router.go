package router

import (
	"encoding/json"
)

type Router struct {
	Route string      // 路由
	Body  interface{} // 包内容
}

// Handler defines route handler, middleware handler type.
type Handler func(context *Context, tag string)

var cacheRouter map[string][]Handler

func init() {
	cacheRouter = make(map[string][]Handler)
	cacheRouter[ONE_TO_MANY] = []Handler{group_chat}
}

func NewRouter(data []byte, tag string) error {
	rt := new(Router)
	err := json.Unmarshal(data, &rt)
	if err != nil {
		return err
	}
	if rs, ok := cacheRouter[rt.Route]; ok {
		c := &Context{Body: rt.Body.(string)}
		for _, r := range rs {
			r(c, tag)
		}
	}
	return nil
}
