package main

import (
	"fmt"
	"github.com/yagou/heiyo/server"
)

func main() {
	server.Websockets()
}

func tcp() {
	fmt.Println("start tcp")
	app := &server.App{}
	app.Run(":5498")
}
