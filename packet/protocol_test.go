package packet

import (
	"testing"
)

func TestPacket(t *testing.T) {
	body1 := Packet([]byte("吴赐有"))
	t.Log(string(body1))
}

func TestUnPacket(t *testing.T) {
	body1 := Packet([]byte("吴赐有"))

	// 声明一个临时缓存
	tempBuffer := make([]byte, 0)

	// 声明一个数据管道用于接收解包后的数据
	readerBody := make(chan []byte, 16)

	tempBuffer = UnPacket(body1[0:3], readerBody)
	t.Logf("tempBuffer: %s \n", tempBuffer)
	tempBuffer = UnPacket(append(tempBuffer, body1[3:]...), readerBody)
	t.Logf("readerBody: %s \n", string(<-readerBody))
	t.Logf("tempBuffer: %s \n", tempBuffer)
}
