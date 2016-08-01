package main

import (
	"fmt"
	"github.com/gernest/utron"
	_ "github.com/yagou/heiyo/controller"
	"github.com/yagou/heiyo/server"
)

func main() {

	fmt.Println("hello")
	fmt.Println("hello1")
	// utron_web()
}

func utron_web() {
	utron.Run()
}

func websocket() {
	server.Websockets()
}

func tcp() {
	fmt.Println("start tcp")
	app := &server.App{}
	app.Run(":5498")
}
