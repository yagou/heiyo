package main

import (
	"../packet"
	"fmt"
	"net"
)

var connList map[int]net.Conn

func main() {
	connList = make(map[int]net.Conn)
	listen, err := net.Listen("tcp", ":5498")
	if err != nil {
		panic(err)
	}
	var connLenth = 2
	var connTag = 0
	for ; ; connTag++ {
		conn, err := listen.Accept()
		currentTag := connTag
		if err != nil {
			continue
		}
		if len(connList) <= connLenth {
			connList[connTag] = conn
		} else {
			conn.Write(packet.Packet([]byte("当前连接数太多，请稍后再试")))
			conn.Close()
			continue
		}

		go func() {
			buffer := make([]byte, 1024)
			tempbuffer := make([]byte, 0)
			readerBody := make(chan []byte, 2)

			go reader(readerBody, currentTag)

			// 读取数据
			for {
				n, err := conn.Read(buffer)
				if err != nil || n <= 0 {
					delete(connList, currentTag)
					break
				}
				tempbuffer = packet.UnPacket(append(tempbuffer, buffer[:n]...), readerBody)
			}

		}()

	}
	fmt.Println("hello word server")
}

func reader(readerBody chan []byte, currentTag int) {
	for {
		select {
		case body := <-readerBody:
			fmt.Println(string(body))
			for k, _ := range connList {
				if k == currentTag {
					continue
				}
				connList[k].Write(packet.Packet(body))
			}
		}
	}
}
