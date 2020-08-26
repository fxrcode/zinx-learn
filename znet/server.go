package znet

import (
	"fmt"
	"net"
	"time"
	"zinx/ziface"
)

// Server is impl of IServer
type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

// Start impl IServer
func (s *Server) Start() {
	fmt.Printf("[START] Server listenner at IP: %s, Port: %d, is starting\n", s.IP, s.Port)

	go func() {
		// 1. get one TCP Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("failed to resolve tcp addr, err: ", err)
			return
		}

		// 2. listen to server addr
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("Failed to listen", s.IPVersion, "err", err)
			return
		}

		// succeed to listen addr
		fmt.Println("Start Zinx server ", s.Name, " succeed, now listening...")

		// 3. Start server network connection service
		for {
			// 3.1 Block waiting for client's connection request
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("failed to accept, err: ", err)
				continue
			}

			// 3.2 TODO: Server.Start() configure server max connection, close this new connection once greater than max

			// 3.3 TODO: Server.Start() handle new connection request, and have handler-conn bound

			// fake max 512 char pong
			go func() {
				// polling pings from client
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err ", err)
						continue
					}
					// pong
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("pong back buf err ", err)
						continue
					}
				}
			}()
		}
	}()
}

// Stop impl IServer
func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server, name ", s.Name)
	// TODO: Server.Stop() should clean up connection info and misc ones
}

// Serve impl IServer
func (s *Server) Serve() {
	s.Start()

	// TODO: Server.Serve() handle others when bootstrap service

	// Block, otherwice: Main Go return, listener's Go return
	for {
		time.Sleep(10 * time.Second)
	}
}

// NewServer to reate a handler of server
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      7777,
	}
	return s
}
