package main

import (
	"fmt"
	"net"
	// "time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:5498")
	if err != nil {
		panic(err)
	}
	sms := make([]byte, 128)
	fmt.Print("请输入要发送的信息：")
	fmt.Scan(&sms)
	conn.Write(sms)

}
