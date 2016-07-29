package router

type Context struct {
	// 事件
	Event string
	// 正文
	Body string
	// 房间tag
	Target string
}
