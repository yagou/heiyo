package packet

import (
	"bytes"
	"encoding/binary"
)

const (
	// 协议包头
	ConstHanger = "ayou_heiyo"
	// 包头长度
	ConstHangerLength = len(ConstHanger)
	// 数据包长度
	ConstSaveDataLength = 4
)

// 装包
func Packet(body []byte) []byte {
	return append(append([]byte(ConstHanger), IntToBytes(len(body))...), body...)
}

func UnPacket(buffer []byte, readerChannel chan []byte) []byte {
	length := len(buffer)

	var i int
	for i = 0; i < length; i++ {
		// 缓存数据长度小于 (包头长度+数据标识长度) 直接跳出
		if length < ConstHangerLength+ConstSaveDataLength {
			break
		}

		if string(buffer[i:i+ConstHangerLength]) == ConstHanger {
			// 获取数据包的实际长度
			messageLength := ByteToInt(buffer[i+ConstHangerLength : i+ConstHangerLength+ConstSaveDataLength])

			// 当前缓存数据长度小于包长度
			if length < i+ConstHangerLength+ConstSaveDataLength+messageLength {
				break
			}
			data := buffer[i+ConstHangerLength+ConstSaveDataLength : i+ConstHangerLength+ConstSaveDataLength+messageLength]
			readerChannel <- data
			i += ConstHangerLength + ConstSaveDataLength + messageLength - 1
		}
	}
	if i == length {
		return make([]byte, 0)
	}
	return buffer[i:]
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