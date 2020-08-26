package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("Client test ... start")
	// give 3 sec for server to bootstrap
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start failed, err ", err)
		return
	}

	for {
		_, err := conn.Write([]byte("hahaha"))
		if err != nil {
			fmt.Println("write failed, err ", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err ", err)
			return
		}
		fmt.Printf(" server call back: %s, cnt = %d\n", buf, cnt)
		time.Sleep(1 * time.Second)
	}
}
