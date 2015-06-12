package packet

import (
	// "fmt"
	"testing"
)

// func _TestHeader(t *testing.T) {
// 	packetHeader := NewNetPacket()
// 	t.Log(packetHeader.Header())
// }

func TestIntToBytes(t *testing.T) {
	packet := Packet([]byte("你现在还好吗？你现在还好吗？你现在还好吗？你现在还好吗？你现在还好吗？你现在还好吗？你现在还好吗？你现在还好吗？你现在还好吗？你现在还好吗？你现在还好吗？你现在还好吗？你现在还好吗？你现在还好吗？"))
	t.Log(packet)

	netPacket := UnPacket(packet)

	t.Logf("%s", netPacket.DataBody)

}
