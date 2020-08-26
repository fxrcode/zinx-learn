package main

import "zinx/znet"

func main() {
	// 1. create a server handler
	s := znet.NewServer("[zinx v0.1]")
	// 2. start serving
	s.Serve()
}
