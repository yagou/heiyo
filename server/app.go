package server

import (
	"fmt"
	"github.com/yagou/heiyo/conn"
	"github.com/yagou/heiyo/packet"
	"github.com/yagou/heiyo/router"
	"net"
)

type App struct {
}

func (app *App) Run(laddr string) {
	listen, err := net.Listen("tcp", laddr)
	if err != nil {
		panic(err)
	}

	// 获取全局唯一的连接缓存
	hyc := conn.NewHYConn()
	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}

		tag, err := hyc.Add_conn(conn)

		if err != nil {

			continue
		}

		go func() {
			buffer := make([]byte, 1024)
			tempbuffer := make([]byte, 0)
			readerBody := make(chan []byte, 2)

			go func(readerBody chan []byte, tag string) {
				for {
					select {
					case body := <-readerBody:
						fmt.Println(body)
						go router.NewRouter(body, tag)
					}
				}
			}(readerBody, tag)

			// 读取数据
			for {
				n, err := conn.Read(buffer)
				if err != nil || n <= 0 {
					break
				}

				tempbuffer = packet.UnPacket(append(tempbuffer, buffer[:n]...), readerBody)
			}

		}()

	}
}
