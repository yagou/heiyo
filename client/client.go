package main

import (
	"../packet"
	"fmt"
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
		// fmt.Print("请输入要发送的信息：")
		fmt.Scan(&sms)
		if string(sms) == "exit" {
			fmt.Println("bye ...")
			break
		}
		conn.Write(packet.Packet(sms))
	}
}
