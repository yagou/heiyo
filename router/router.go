package router

import (
	"encoding/json"
	"fmt"
)

type Router struct {
	Route   string // 路由
	Context *Context
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
	fmt.Println(string(data))
	err := json.Unmarshal(data, &rt)
	fmt.Println(rt)
	if err != nil {
		return err
	}
	if rs, ok := cacheRouter[rt.Route]; ok {
		for _, r := range rs {
			r(rt.Context, tag)
		}
	}
	return nil
}
