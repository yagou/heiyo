package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":5498")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		go func() {
			buffer := make([]byte, 2)
			var body []byte
			for {
				fmt.Println("start read ...")
				lens, err := conn.Read(buffer)
				fmt.Println("start end ...")
				fmt.Println(err)
				if err != nil || lens <= 0 {
					break
				}
				body = append(body, buffer[:lens]...)
				fmt.Println(lens)
			}
			fmt.Printf("内容是: %s \n ", body)
		}()

	}
	fmt.Println("hello word server")
}
