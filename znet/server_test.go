package znet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func ClientTest() {
	fmt.Println("Client Test ... start")
	// give sometime for Server to bootstraap, so test after 3 sec
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("Client stary err, exit!")
		return
	}

	for {
		_, err := conn.Write([]byte("hello ZINX!"))
		if err != nil {
			fmt.Println("write failed, err: ", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf failed ")
			return
		}

		fmt.Printf(" server call back: %s, cnt = %d\n", buf, cnt)

		time.Sleep(1 * time.Second)
	}
}

func TestServer(t *testing.T) {
	/*
		Test server
	*/
	// 1. create a handler for server
	s := NewServer("zinx V0.1")

	/*
		Test client
	*/
	go ClientTest()

	// 2. start serving
	s.Serve()
}
