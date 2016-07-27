package router

import (
	"testing"
)

func TestNewRouter(t *testing.T) {
	var body = []byte(`
			{"Route":"ONE_TO_ONE","Body":"wuciyou"}
		`)
	//
	//{"Route":"ONE_TO_MANY","Body":"吴赐有，你好啊"}
	//
	err := NewRouter(body)
	t.Log(err)
}
