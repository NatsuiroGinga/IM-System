package main

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	IP   string
	Port int
}

func NewServer(ip string, port int) *Server {
	s := &Server{ip, port}
	return s
}

func (s *Server) Handler(conn net.Conn) {
	fmt.Println("链接建立成功")
}

func (s *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		log.Fatal("net.Listen err: ", err)
	}

	defer func(listener net.Listener) {
		_ = listener.Close()
	}(listener)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			continue
		}
		// do handler
		go s.Handler(conn)
	}
}
