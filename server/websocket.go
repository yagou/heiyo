package server

import (
	"fmt"
	"github.com/yagou/heiyo/conn"
	"github.com/yagou/heiyo/packet"
	"github.com/yagou/heiyo/router"
	"golang.org/x/net/websocket"
	"html/template"
	"net/http"
)

func Websockets() {
	fmt.Println("start websockets")

	http.Handle("/css/", http.FileServer(http.Dir("client")))
	http.Handle("/js/", http.FileServer(http.Dir("client")))
	http.Handle("/fonts/", http.FileServer(http.Dir("client")))

	http.Handle("/images/", http.FileServer(http.Dir("client")))
	http.Handle("/lib/", http.FileServer(http.Dir("client")))

	http.Handle("/ws", websocket.Handler(countServer))
	http.HandleFunc("/", Templates)
	http.ListenAndServe(":5498", nil)
}

func Templates(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("client/client.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func countServer(ws *websocket.Conn) {
	tag, err := conn.NewHYConn().Add_conn(ws)
	defer func() {
		// 关闭连接
		conn.NewHYConn().Del_conn(tag)
	}()

	if err != nil {
		return
	}

	buffer := make([]byte, 3)
	tempbuffer := make([]byte, 0)
	readerBody := make(chan []byte, 2)
	// 读取数据
	go func(readerBody chan []byte, tag string) {
		for {
			select {
			case body := <-readerBody:
				go router.NewRouter(body, tag)
			}
		}
	}(readerBody, tag)

	// 登录誰

	// 读取数据
	for {
		n, err := ws.Read(buffer)
		if err != nil || n <= 0 {
			break
		}
		tempbuffer = packet.UnPacket(append(tempbuffer, buffer[:n]...), readerBody)
	}

}
