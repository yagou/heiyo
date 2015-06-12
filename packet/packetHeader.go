package packet

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// 网络数据包头
type netPacketHeader struct {
	header       string
	AyouDataSize int
}

// 网络包
type netPacket struct {
	PacketHeader netPacketHeader
	DataBody     []byte // 包数据
}

// func NewNetPacket() *netPacket {
// 	return _netPacket
// }
var _netPacket *netPacket

func init() {
	_netPacket = &netPacket{PacketHeader: netPacketHeader{header: "ayou_heiyo"}}
}

func (p *netPacket) Header() string {
	return p.PacketHeader.header
}

// 装包
func Packet(body []byte) []byte {

	return append(append([]byte(_netPacket.PacketHeader.header), IntToBytes(len(body))...), body...)
}

// 解包
func UnPacket(packet []byte) *netPacket {
	var _readeNetPacket = &netPacket{}
	var headerLength = len([]byte(_netPacket.PacketHeader.header))
	fmt.Printf("headerLength:%d \n ", headerLength)
	_readeNetPacket.PacketHeader.header = string(packet[:headerLength])
	_readeNetPacket.PacketHeader.AyouDataSize = ByteToInt(packet[headerLength : headerLength+4])
	_readeNetPacket.DataBody = packet[headerLength+4:]
	return _readeNetPacket
}

// 将数字转换成byte类型
func IntToBytes(length int) []byte {
	x := int32(length)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

// 将byte转换成数字
func ByteToInt(b []byte) int {
	var x int32
	bytesBuffer := bytes.NewBuffer(b)
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)

}
