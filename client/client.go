package main

import (
	"encoding/json"
	"fmt"
	"github.com/yagou/heiyo"
	"github.com/yagou/heiyo/packet"
	"github.com/yagou/heiyo/router"
	"net"
)

func main() {
	client()
}

func client() {
	conn, err := net.Dial("tcp", "127.0.0.1:5498")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	go func() {
		buffer := make([]byte, 1024)
		tempbuffer := make([]byte, 0)
		readerBody := make(chan []byte, 2)
		go func(readerBody chan []byte) {
			for {
				select {
				case body := <-readerBody:
					fmt.Println(string(body))

				}
			}
		}(readerBody)

		// 读取数据
		for {
			n, err := conn.Read(buffer)
			if err != nil || n <= 0 {
				break
			}

			tempbuffer = packet.UnPacket(append(tempbuffer, buffer[:n]...), readerBody)
		}
	}()

	for {
		sms := make([]byte, 128)
		var bodys []byte
		fmt.Println(`
			请选择操作：
			1、一对一聊天
			2、一对多聊天
			3、获取在线列表
		`)
		fmt.Scan(&sms)
		switch string(sms) {
		case "exit":
			fmt.Println("bye ...")
			break
		case "1":
			fmt.Println("暂时不支持一对一聊天")
		case "2":
			fmt.Println("请输入内容")
			fmt.Scan(&sms)

			rt := new(heiyo.Router)

			rt.Route = router.ONE_TO_MANY
			rt.Body = string(sms)
			bodys, _ = json.Marshal(rt)
		case "3":
			rt := new(heiyo.Router)

			rt.Route = router.GET_USER_LIST
			rt.Body = string(sms)
			bodys, _ = json.Marshal(rt)
		}

		wstring := packet.Packet(bodys)
		fmt.Printf("发送的是:%s \n", string(wstring))
		conn.Write(wstring)
	}
}
