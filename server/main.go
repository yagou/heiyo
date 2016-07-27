package main

import (
	"fmt"
	"github.com/yagou/heiyo"
	"golang.org/x/net/websocket"
	"net/http"
	"time"
)

func main() {
	websockets()
}

type Count struct {
	S string
	N int
}

func websockets() {
	fmt.Println("start websockets")
	http.Handle("/ws", websocket.Handler(countServer))
	http.ListenAndServe(":5498", nil)
}

func countServer(ws *websocket.Conn) {
	fmt.Println("count")
	var counts = 0
	// 读取数据
	go func() {
		buffer := make([]byte, 1024)

		// 读取数据
		for {
			n, err := ws.Read(buffer)
			if err != nil || n <= 0 {
				break
			}
			fmt.Printf("接收到的数据是：%s \n", buffer)
		}

	}()

	for {
		ws.Write([]byte(fmt.Sprintf("连接成功： %d \n", counts)))
		counts++
		time.Sleep(time.Second)
	}
}

func tcp() {
	fmt.Println("start tcp")
	app := &heiyo.App{}
	app.Run(":5498")
}
