package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.1.0.1:5498")
	if err != nil {
		panic(err)
	}
	// for {
	body := fmt.Sprintf(" 当前是：%d 秒 ", time.Now().Second())
	fmt.Println(body)
	conn.Write([]byte("测试的1"))
	time.Sleep(time.Second)
	// }

}
